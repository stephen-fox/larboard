package jsonw

import (
	"encoding/json"
)

func ToString(target interface{}) (string, error) {
	raw, err := json.Marshal(target)
	if err != nil {
		return "", err
	}

	return string(raw), nil
}

func StringToStruct(jsonStr string, pointerToObject interface{}) error {
	return json.Unmarshal([]byte(jsonStr), pointerToObject)
}

func ToPrettyString(object interface{}) ([]byte, error) {
	raw, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		return []byte("\n"), err
	}

	return raw, nil
}