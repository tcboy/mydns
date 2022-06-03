package util_json

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
)

var (
	NewDecoder = jsoniter.NewDecoder
	NewEncoder = jsoniter.NewEncoder
)

// still under dev:test
// !!use with caution!!
func DefaultMarshal(v interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	e := jsoniter.NewEncoder(buffer)
	e.SetEscapeHTML(false)
	err := e.Encode(v)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func Marshal(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}

func DefaultUnmarshal(data []byte, v interface{}) error {
	d := jsoniter.NewDecoder(bytes.NewBuffer(data))
	d.UseNumber()
	return d.Decode(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return jsoniter.Unmarshal(data, v)
}
