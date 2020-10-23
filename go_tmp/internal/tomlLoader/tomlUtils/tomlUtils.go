package tomlUtils

import "fmt"

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
	}
	for _, w := range typeToMatch {
		if w == t {
			return true
		}
	}
	return false
}
