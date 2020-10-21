package utils

import (
	"io/ioutil"
	"log"
)

func IsIn(str string, arr []string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func GetFileContent(fp string) string {
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
