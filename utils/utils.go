package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"html"
	"html/template"
	"math/rand"
	"mydns/utils/util_json"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GenerateLogId() string {
	var uuidStr string
	newUUID := uuid.NewV1()

	uuidStr = newUUID.String()
	uuidStr = strings.Replace(uuidStr, "-", "s", -1)

	return uuidStr
}

func Sha1(data string) string {
	s := sha1.New()
	s.Write([]byte(data))
	return hex.EncodeToString(s.Sum([]byte(nil)))
}

func ArrayToString(slice interface{}, separator string) string {
	array := TransInterfacToSlice(slice)
	res := ""
	for i, val := range array {
		if i != 0 {
			res += separator
		}
		res += ToString(val)
	}
	return res
}

func BoolToInt(b bool) int8 {
	if b {
		return int8(1)
	}
	return int8(0)
}

func CopyInterface(src, dst interface{}) error {
	dataByte, err := util_json.Marshal(src)
	if err != nil {
		return err
	}
	err = Unmarshal(string(dataByte), &dst)
	if err != nil {
		return err
	}
	return nil
}

func EqualsTo(a interface{}) func(interface{}) bool {
	return func(b interface{}) bool {
		return a == b
	}
}

func NotEqualsTo(a interface{}) func(interface{}) bool {
	return func(b interface{}) bool {
		return a != b
	}
}

func GreaterThan(a interface{}) func(interface{}) bool {
	return func(b interface{}) bool {
		return b.(int64) > a.(int64)
	}
}

func LessThan(a interface{}) func(interface{}) bool {
	return func(b interface{}) bool {
		return a.(int64) > b.(int64)
	}
}

func FilterDupInt64ByArrayB(arrA []int64, arrB []int64) []int64 {
	if len(arrA) == 0 || len(arrB) == 0 {
		return arrA
	}

	res := []int64{}
	sort.Slice(arrA, func(i, j int) bool { return arrA[i] < arrA[j] })
	sort.Slice(arrB, func(i, j int) bool { return arrB[i] < arrB[j] })
	j := 0
	for i := 0; i < len(arrA); i++ {
		for ; j < len(arrB); j++ {
			if arrB[j] >= arrA[i] {
				break
			}
		}
		if j == len(arrB) || arrB[j] > arrA[i] {
			res = append(res, arrA[i])
		}
	}

	return res
}

func FilterDuplicates(elems interface{}) []interface{} {
	elemIntrs := TransInterfacToSlice(elems)
	encountered := make(map[interface{}]bool)

	j := 0
	for i := 0; i < len(elemIntrs); i++ {
		if _, found := encountered[elemIntrs[i]]; !found {
			if i != j {
				elemIntrs[j] = elemIntrs[i]
			}
			j++
			encountered[elemIntrs[i]] = true
		}
	}
	return elemIntrs[:j]
}

func FilterInt32(a []int32, lambda func(interface{}) bool) (ret []int32) {
	for _, ele := range a {
		if lambda(ele) {
			ret = append(ret, ele)
		}
	}
	return
}

func FilterInt64(a []int64, lambda func(interface{}) bool) (ret []int64) {
	for _, ele := range a {
		if lambda(ele) {
			ret = append(ret, ele)
		}
	}
	return
}

func FilterString(a []string, lambda func(interface{}) bool) (ret []string) {
	for _, ele := range a {
		if lambda(ele) {
			ret = append(ret, ele)
		}
	}
	return
}

func StringEquals(a string, b string) bool {
	return strings.Compare(a, b) == 0
}

func GetFromNil(i interface{}, t reflect.Type) reflect.Value {
	if i != nil {
		return reflect.ValueOf(i)
	}
	return reflect.Zero(t)
}

func GenMD5(s string, salt string) string {
	if salt != "" {
		s = GetStrFromTpl("%v_%v", s, salt)
	}
	h := md5.New()
	h.Write([]byte(s))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func GetStrFromTpl(tpl string, args ...interface{}) string {
	return fmt.Sprintf(tpl, args...)
}

func InterfaceConv(i interface{}, defaultValue interface{}, valueType string) interface{} {
	stringValue := ToString(i)
	return StringConv(stringValue, defaultValue, valueType)
}

func JsonMarshalOutput(d interface{}) string {
	buffer := &bytes.Buffer{}
	encoder := util_json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(d); err != nil {
		return ""
	}
	return string(buffer.Bytes())
}

func MapInt32Keys(imap interface{}) []int32 {
	values := reflect.ValueOf(imap).MapKeys()
	keys := make([]int32, 0, len(values))
	for _, v := range values {
		keys = append(keys, v.Interface().(int32))
	}
	return keys
}

func MapInt64Keys(imap interface{}) []int64 {
	values := reflect.ValueOf(imap).MapKeys()
	keys := make([]int64, 0, len(values))
	for _, v := range values {
		keys = append(keys, v.Interface().(int64))
	}
	return keys
}

func MapStringKeys(imap interface{}) []string {
	values := reflect.ValueOf(imap).MapKeys()
	keys := make([]string, 0, len(values))
	for _, v := range values {
		keys = append(keys, v.Interface().(string))
	}
	return keys
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func MinInt16(x, y int16) int16 {
	if x < y {
		return x
	}
	return y
}

func MaxInt16(x, y int16) int16 {
	if x > y {
		return x
	}
	return y
}

func MinInt32(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func MaxInt32(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func MergeJson(src, dst string) (string, error) {
	if len(src) < 3 {
		return dst, nil
	}

	from, to := make(map[string]interface{}), make(map[string]interface{})
	Unmarshal(src, &from)
	Unmarshal(dst, &to)
	for k, v := range from {
		to[k] = v
	}
	data, err := util_json.Marshal(&to)
	return string(data), err
}

func MergeJsonStrFromMap(srcStr string, data map[string]interface{}) (string, error) {
	src := make(map[string]interface{})
	Unmarshal(srcStr, &src)

	for k, v := range data {
		src[k] = v
	}

	rawData, err := util_json.Marshal(&src)
	return string(rawData), err
}

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

func StringToMap(str string) map[string]interface{} {
	m := map[string]interface{}{}
	if err := Unmarshal(str, &m); err != nil {
		return nil
	}
	return m
}

func ToError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToJsonString(i interface{}) string {
	if i == nil {
		return ""
	}

	data, err := util_json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(data)
}

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

func TransInterfacToSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("TransInterfacToSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func UnescapedString(unescaped string) string {
	return html.UnescapeString(unescaped)
}

func Unmarshal(data string, v interface{}) error {
	d := util_json.NewDecoder(strings.NewReader(data))
	d.UseNumber()
	return d.Decode(v)
}

func UrlEncoded(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		return ""
	}
	return u.String()
}

func SliceInt8Contains(slice []int8, element int8) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func SliceInt16Contains(slice []int16, element int16) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func SliceInt32Contains(slice []int32, element int32) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func SliceInt64Contains(slice []int64, element int64) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func SliceIntrContains(slice interface{}, ele interface{}) bool {
	for _, e := range TransInterfacToSlice(slice) {
		if e == ele {
			return true
		}
	}
	return false
}

func SliceStrContains(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func StringIndexBy(i int, s []string) string {
	if i < 0 || i >= len(s) {
		return ""
	}
	return s[i]
}

// Returns an int >= min, < max
func RandomIntRange(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

func BoolToIntStr(v bool) string {
	if v {
		return "1"
	} else {
		return "0"
	}
}
