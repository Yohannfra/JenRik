package utils

import (
	"fmt"
	"github.com/Yohannfra/JenRik/internal/logLevel"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
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

func FileExists(fp string) bool {
	if _, err := os.Stat(fp); !os.IsNotExist(err) {
		return true
	}
	return false
}

func PrintDiff(gotStr string, expectedStr string) {
	if logLevel.LOG_LEVEL == logLevel.QUIET {
		return
	}
	maxLen := math.Max(float64(len(gotStr)), float64(len(expectedStr)))
	fmt.Println(strings.Repeat("-", int(math.Min(30, maxLen))))
	fmt.Println("Expected:'")
	fmt.Print(gotStr)
	fmt.Print("'\n")
	fmt.Println("Bug got:'")
	fmt.Print(expectedStr)
	fmt.Print("'\n")
	fmt.Println(strings.Repeat("-", int(math.Min(30, maxLen))))
}

func CompareStrArray(arr1 []string, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, _ := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
