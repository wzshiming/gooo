package gooo

import "encoding/json"

func Decode(b []byte, s interface{}) error {
	return json.Unmarshal(b, s)
}

func Encode(s interface{}) ([]byte, error) {
	return json.Marshal(s)
}
