package util

import (
	"encoding/json"
	"sort"
)

func IndexOfString(target string, arr []string) int {
	return sort.SearchStrings(arr, target)
}
func Map2Obj(m interface{}, obj interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, obj)
}
func Obj2Str(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
