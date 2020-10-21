package utils

import (
	"log"
	"io/ioutil"
)

func Is_in(str string, arr []string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func Get_file_content(fp string) string {
    content, err := ioutil.ReadFile(fp)
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}
