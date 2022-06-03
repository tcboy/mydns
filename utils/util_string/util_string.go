package util_string

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

func ToString(i interface{}) string {
	if i == nil {
		return ""
	}
	switch s := i.(type) {
	case string:
		return s
	case bool:
		return strconv.FormatBool(s)
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32)
	case int:
		return strconv.Itoa(s)
	case int64:
		return strconv.FormatInt(s, 10)
	case int32:
		return strconv.Itoa(int(s))
	case int16:
		return strconv.FormatInt(int64(s), 10)
	case int8:
		return strconv.FormatInt(int64(s), 10)
	case uint:
		return strconv.FormatInt(int64(s), 10)
	case uint64:
		return strconv.FormatInt(int64(s), 10)
	case uint32:
		return strconv.FormatInt(int64(s), 10)
	case uint16:
		return strconv.FormatInt(int64(s), 10)
	case uint8:
		return strconv.FormatInt(int64(s), 10)
	case []byte:
		return string(s)
	case template.HTML:
		return string(s)
	case template.URL:
		return string(s)
	case template.JS:
		return string(s)
	case template.CSS:
		return string(s)
	case template.HTMLAttr:
		return string(s)
	case nil:
		return ""
	case fmt.Stringer:
		return s.String()
	case error:
		return s.Error()
	default:
		return fmt.Sprint(i)
	}
}

//StringConv将字符串转化为其他类型
func StringConv(s string, defaultValue interface{}, valueType string) interface{} {
	var err error
	var value interface{}
	switch valueType {
	case "bool":
		value, err = strconv.ParseBool(s)
	case "int":
		value, err = strconv.ParseInt(s, 10, 0)
		value = int(value.(int64))
	case "int8":
		value, err = strconv.ParseInt(s, 10, 8)
		value = int8(value.(int64))
	case "int16":
		value, err = strconv.ParseInt(s, 10, 16)
		value = int16(value.(int64))
	case "int32":
		value, err = strconv.ParseInt(s, 10, 32)
		value = int32(value.(int64))
	case "int64":
		value, err = strconv.ParseInt(s, 10, 64)
	case "float64":
		value, err = strconv.ParseFloat(s, 64)
	case "string":
		value = s
		if value == "" {
			value = defaultValue.(string)
		}
	case "[]string":
		value = strings.Fields(s)
	default:
		panic(fmt.Sprintf("nonexistent value type: %s", valueType))
	}
	if err != nil {
		value = defaultValue
	}
	return value
}

func StringPtrIfNotEmpty(input string) *string {
	if input == "" {
		return nil
	}

	return &input
}

func StringToArray(str string, separator string) []interface{} {
	array := strings.Split(str, separator)
	res := make([]interface{}, len(array))
	for i, val := range array {
		res[i] = val
	}
	return res
}

func SplitToInt64Array(input string) []int64 {
	ary := strings.Split(input, ",")
	ret := make([]int64, 0)
	for _, s := range ary {
		if len(s) > 0 {
			elm, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				continue
			}
			ret = append(ret, elm)
		}
	}
	return ret
}

func SplitToInt32Array(input string) []int32 {
	ary := strings.Split(input, ",")
	ret := make([]int32, 0)
	for _, s := range ary {
		elm, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			continue
		}
		ret = append(ret, int32(elm))
	}
	return ret
}
