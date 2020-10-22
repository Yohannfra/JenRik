package parser

import (
	"fmt"
	"github.com/Yohannfra/JenRik/internal/tester"
	"github.com/Yohannfra/JenRik/internal/utils"
	"github.com/pelletier/go-toml"
	"log"
	"os"
	"os/exec"
	"strings"
)

func isStringArray(data interface{}) bool {
	arr := data.([]interface{})

	for _, d := range arr {
		if !isOfType(d, []string{"string"}) {
			return false
		}
	}
	return true
}

func isOfType(data interface{}, typeToMatch []string) bool {
	t := fmt.Sprintf("%T", data)

	if t == "[]interface {}" { // check for str array
		return isStringArray(data)
	}
	for _, w := range typeToMatch {
		if w == t {
			return true
		}
	}
	return false
}

func CheckTestsValidity(testName string, testDict *toml.Tree) {
	testsKeysConfig := []struct {
		name             string
		types            []string
		incompatibleWith string
		required         bool
	}{
		{"args", []string{"strList"}, "", true},
		{"status", []string{"int64"}, "", true},
		{"stdout", []string{"strList", "string"}, "stdout_file", false},
		{"stderr", []string{"strList", "string"}, "stderr_file", false},
		{"pre", []string{"strList", "string"}, "", false},
		{"post", []string{"strList", "string"}, "", false},
		{"stdout_file", []string{"string"}, "stdout", false},
		{"stderr_file", []string{"string"}, "stderr", false},
		{"pipe_stdout", []string{"string"}, "", false},
		{"pipe_stderr", []string{"string"}, "", false},
		{"timeout", []string{"int64"}, "", false},
		{"should_fail", []string{"bool"}, "", false},
		{"stdin", []string{"strList", "string"}, "stdin_file", false},
		{"stdin_file", []string{"string"}, "stdin", false},
		{"env", []string{"dict"}, "", false},
		{"add_env", []string{"dict"}, "", false},
		{"repeat", []string{"int64"}, "", false},
	}

	for _, key := range testDict.Keys() {
		found := false
		var types []string

		for _, keyConfig := range testsKeysConfig {
			if keyConfig.name == key {
				found = true
				types = keyConfig.types
			}
			if key == keyConfig.name && utils.IsIn(keyConfig.incompatibleWith, testDict.Keys()) {
				log.Fatalf("%s: Incompatible keys, %s and %s", testName, keyConfig.name, keyConfig.incompatibleWith)
			}

			if keyConfig.required && !utils.IsIn(keyConfig.name, testDict.Keys()) {
				log.Fatalf("%s : Missing field : %s", testName, keyConfig.name)
			}
		}
		if found == false {
			log.Fatalf("Unknown key: %s", key)
		}
		value := testDict.Get(key)
		if !isOfType(value, types) {
			log.Fatalf("Error in test '%s' : key '%s' must be '%v' but it's '%T'", testName, key, types, value)
		}
	}
}

func runBuildCommand(command string) {
	tmp := strings.Split(command, " ")
	var cmd *exec.Cmd

	if len(tmp) == 1 {
		cmd = exec.Command(tmp[0])
	} else {
		cmd = exec.Command(tmp[0], strings.Join(tmp[1:], " "))
	}

	err := cmd.Run()
	if err != nil {
		log.Fatalln("Error running build command", err)
	}
}

func fixBinaryPath(binaryPath string, fp string) string {
	splittedPath := strings.Split(fp, "/")
	splittedPath = splittedPath[:len(splittedPath)-1]

	if len(splittedPath) == 0 { // from the same directory, nothing to do
		return binaryPath
	}
	pathToToml := strings.Join(splittedPath, "/") + "/"
	if pathToToml == "/" && !strings.Contains(fp, "/") {
		pathToToml = "./"
	}
	binaryPath = pathToToml + binaryPath
	binaryPath = strings.ReplaceAll(binaryPath, "././", "./")
	binaryPath = strings.ReplaceAll(binaryPath, "/./", "/")
	return binaryPath
}

func checkBinaryValidity(fp string) {
	if !utils.FileExists(fp) {
		log.Fatalf("%s : file not found\n", fp)
	}
	fi, err := os.Stat(fp)
	if err != nil {
		log.Fatal(err)
	}
	if fi.IsDir() {
		log.Fatalf("%s : is a directory\n", fp)
	}
	if !(fi.Mode().Perm()&0111 != 0) { // if file is not executable
		log.Fatalf("%s : is not executable\n", fp)
	}
}

func LoadTestFile(fp string) tester.TestSuiteData {
	var testData tester.TestSuiteData
	var err error

	fc := utils.GetFileContent(fp)
	testData.TomlContent, err = toml.Load(fc)
	testData.BinaryPath = ""

	if err != nil {
		log.Fatal(err)
	}
	for _, key := range testData.TomlContent.Keys() {
		if key == "binary_path" {
			testData.BinaryPath = testData.TomlContent.Get(key).(string)
		} else if key == "build_command" {
			runBuildCommand(testData.TomlContent.Get(key).(string))
		} else {
			CheckTestsValidity(key, testData.TomlContent.Get(key).(*toml.Tree))
		}
	}
	if testData.BinaryPath == "" {
		log.Fatal("Could not find binary_path key in", fp)
	}

	testData.BinaryPath = fixBinaryPath(testData.BinaryPath, fp)
	checkBinaryValidity(testData.BinaryPath)
	return testData
}
