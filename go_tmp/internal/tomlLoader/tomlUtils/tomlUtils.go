package tomlUtils

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

func IsStringMap(data *toml.Tree) bool {
	for _, d := range data.Keys() {
		if !IsOfType(d, []string{"string"}) || !IsOfType(data.Get(d), []string{"string"}) {
			return false
		}
	}
	return true
}

func IsStringArray(data interface{}) bool {
	arr := data.([]interface{})

	for _, d := range arr {
		if !IsOfType(d, []string{"string"}) {
			return false
		}
	}
	return true
}

func IsOfType(data interface{}, typeToMatch []string) bool {
	t := fmt.Sprintf("%T", data)

	if t == "[]interface {}" { // check for str array
		return IsStringArray(data)
	} else if t == "*toml.Tree" { // for env
		return IsStringMap(data.(*toml.Tree))
	}
	for _, w := range typeToMatch {
		if w == t {
			return true
		}
	}
	return false
}

func ToStrArr(data interface{}) []string {
	var arr []string
	tmp := data.([]interface{})

	for _, elem := range tmp {
		arr = append(arr, elem.(string))
	}
	return arr
}

func ToStrMap(data *toml.Tree) map[string]string {
	var dict map[string]string
	dict = make(map[string]string)

	for _, d := range data.Keys() {
		dict[d] = data.Get(d).(string)
	}
	return dict
}
