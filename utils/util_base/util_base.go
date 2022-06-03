package util_base

import (
	"fmt"
)

//DefaultValue读取任何类型做type assert，错误时，返回默认值
func DefaultValue(v interface{}, defaultValue interface{}, valueType string) interface{} {
	var (
		ok    bool
		value interface{}
	)
	switch valueType {
	case "bool":
		value, ok = v.(bool)
	case "int":
		value, ok = v.(int)
	case "int8":
		value, ok = v.(int8)
	case "int16":
		value, ok = v.(int16)
	case "int32":
		value, ok = v.(int32)
	case "int64":
		value, ok = v.(int64)
	case "float64":
		value, ok = v.(float64)
	case "string":
		value, ok = v.(string)
	case "[]string":
		value, ok = v.([]string)
	case "[]int32":
		value, ok = v.([]int32)
	case "[]int64":
		value, ok = v.([]int64)
	case "map[string]string":
		value, ok = v.(map[string]string)
	case "map[string]interface{}":
		value, ok = v.(map[string]interface{})
	case "map[int64]interface{}":
		value, ok = v.(map[int64]interface{})
	case "map[int32]interface{}":
		value, ok = v.(map[int32]interface{})
	default:
		panic(fmt.Sprintf("nonexistent value type: %s", valueType))
	}
	if ok {
		return value
	}
	return defaultValue
}
