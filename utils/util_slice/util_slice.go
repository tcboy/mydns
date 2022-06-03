package util_slice

import (
	"encoding/json"
	"strconv"
	"strings"
)

//FilterDupStr 去重string,后重复的将被去掉，后面的数据迁移，保证顺序不变
func FilterDupStr(elems []string) interface{} {
	encountered := make(map[interface{}]bool)
	j := 0
	for i := 0; i < len(elems); i++ {
		if _, found := encountered[elems[i]]; !found {
			if i != j {
				elems[j] = elems[i]
			}
			j++
			encountered[elems[i]] = true
		}
	}
	return elems[:j]
}

//FilterDupInt64 去重int64,后重复的将被去掉，后面的数据迁移，保证顺序不变
func FilterDupInt64(elems []int64) interface{} {
	encountered := make(map[interface{}]bool)
	j := 0
	for i := 0; i < len(elems); i++ {
		if _, found := encountered[elems[i]]; !found {
			if i != j {
				elems[j] = elems[i]
			}
			j++
			encountered[elems[i]] = true
		}
	}
	return elems[:j]
}

func MapElement(data []interface{}, index int64, defaultValue map[string]interface{}) map[string]interface{} {
	if index > int64(len(data))-1 {
		return defaultValue
	}

	v := data[index]
	mapV, ok := v.(map[string]interface{})
	if !ok {
		return defaultValue
	}
	return mapV
}

func StringElement(data []interface{}, index int64, defaultValue string) string {
	if index > int64(len(data))-1 {
		return defaultValue
	}

	v := data[index]
	strV, ok := v.(string)
	if !ok {
		return defaultValue
	}
	return strV
}

func SliceElement(data []interface{}, index int64, defaultValue []interface{}) []interface{} {
	if index > int64(len(data))-1 {
		return defaultValue
	}

	v := data[index]
	sliceV, ok := v.([]interface{})
	if !ok {
		return defaultValue
	}
	return sliceV
}

func Int64Element(data []interface{}, index int64, defaultValue int64) int64 {
	if index > int64(len(data))-1 {
		return defaultValue
	}

	v := data[index]
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

func Float64Element(data []interface{}, index int64, defaultValue float64) float64 {
	if index > int64(len(data))-1 {
		return defaultValue
	}

	v := data[index]
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

//FilterDupInt64 去重int64,后重复的将被去掉，后面的数据迁移，保证顺序不变
func FilterDupInt64ToInt64(elems []int64) []int64 {
	encountered := make(map[int64]bool)
	j := 0
	for i := 0; i < len(elems); i++ {
		if _, found := encountered[elems[i]]; !found {
			if i != j {
				elems[j] = elems[i]
			}
			j++
			encountered[elems[i]] = true
		}
	}
	return elems[:j]
}

func Int64ArrayJoin(ids []int64) string {
	var strs = make([]string, 0, len(ids))
	for _, id := range ids {
		strs = append(strs, strconv.FormatInt(id, 10))
	}
	return strings.Join(strs, ",")
}

func StringSliceDelete(slice []string, target string) []string {
	index := 0
	for _, ele := range slice {
		if ele != target {
			slice[index] = ele
			index++
		}
	}
	return slice[:index]
}
