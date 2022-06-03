package util_map

import (
	"encoding/json"

	"code.aliyun.com/batcn/common/utils/util_string"
)

//StrListToMap 将string转化为Map结构，方便做in查询
func StrListToMap(strList []string) map[string]interface{} {
	strMap := make(map[string]interface{}, len(strList))
	for _, str := range strList {
		strMap[str] = nil
	}

	return strMap
}

//MergeSIMap merge key为String和value为interface的map
//amap 第一个map，在useNewMap等于false的情况下，新的map直接写入第一个map
//bmap 第二个map，相同的key覆盖第一个map内部的值
//useNewMap 是否重新分配一个map返回，为了性能考虑，可以为false，使用第一个map，但需要注意：传入的amap已经被修改了
func MergeSIMap(amap map[string]interface{}, bmap map[string]interface{}, useNewMap bool) map[string]interface{} {
	var newMap map[string]interface{}
	if useNewMap {
		newMap = make(map[string]interface{}, len(amap)+len(bmap))
		for k, v := range amap {
			newMap[k] = v
		}
	}

	for k, v := range bmap {
		if useNewMap {
			newMap[k] = v
		} else {
			amap[k] = v
		}
	}
	if useNewMap {
		return newMap
	}
	return amap
}

func MergeSSSIMap(amap map[string]string, bmap map[string]interface{}, useNewMap bool) map[string]string {
	var newMap map[string]string
	if useNewMap {
		newMap = make(map[string]string, len(amap)+len(bmap))
		for k, v := range amap {
			newMap[k] = v
		}
	}

	for k, v := range bmap {
		if useNewMap {
			newMap[k] = util_string.ToString(v)
		} else {
			amap[k] = util_string.ToString(v)
		}
	}
	if useNewMap {
		return newMap
	}
	return amap
}

func MergeSSMap(amap map[string]string, bmap map[string]string, useNewMap bool) map[string]string {
	var newMap map[string]string
	if useNewMap {
		newMap = make(map[string]string, len(amap)+len(bmap))
		for k, v := range amap {
			newMap[k] = v
		}
	}

	for k, v := range bmap {
		if useNewMap {
			newMap[k] = v
		} else {
			amap[k] = v
		}
	}
	if useNewMap {
		return newMap
	}
	return amap
}

func StringElement(data map[string]interface{}, key string, defaultValue string) string {
	v, ok := data[key]
	if !ok {
		return defaultValue
	}

	strV, ok := v.(string)
	if !ok {
		return defaultValue
	}

	return strV
}

func MapElement(data map[string]interface{}, key string, defaultValue map[string]interface{}) map[string]interface{} {
	v, ok := data[key]
	if !ok {
		return defaultValue
	}

	mapV, ok := v.(map[string]interface{})
	if !ok {
		return defaultValue
	}

	return mapV
}

func SliceElement(data map[string]interface{}, key string, defaultValue []interface{}) []interface{} {
	v, ok := data[key]
	if !ok {
		return defaultValue
	}

	sliceV, ok := v.([]interface{})
	if !ok {
		return defaultValue
	}

	return sliceV
}

func Int64Element(data map[string]interface{}, key string, defaultValue int64) int64 {
	v, ok := data[key]
	if !ok {
		return defaultValue
	}

	switch v.(type) {
	case int:
		return int64(v.(int))
	case int32:
		return int64(v.(int32))
	case int64:
		return v.(int64)
	case json.Number:
		jv, err := v.(json.Number).Int64()
		if err == nil {
			return jv
		} else {
			return defaultValue
		}
	default:
		return defaultValue
	}
}

func Float64Element(data map[string]interface{}, key string, defaultValue float64) float64 {
	v, ok := data[key]
	if !ok {
		return defaultValue
	}

	switch v.(type) {
	case float32:
		return float64(v.(float32))
	case float64:
		return v.(float64)
	case json.Number:
		jv, err := v.(json.Number).Float64()
		if err == nil {
			return jv
		} else {
			return defaultValue
		}
	default:
		return defaultValue
	}
}
