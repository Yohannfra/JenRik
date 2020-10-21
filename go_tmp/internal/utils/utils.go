package utils

import (
	"log"
	"io/ioutil"
)


func is_in(str string, arr []string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func get_file_content(fp string) string {
    content, err := ioutil.ReadFile(fp)
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}
