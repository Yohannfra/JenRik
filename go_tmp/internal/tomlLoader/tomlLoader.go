package tomlLoader

import (
	"fmt"
	"github.com/Yohannfra/JenRik/internal/testData"
	"github.com/Yohannfra/JenRik/internal/tester"
	"github.com/Yohannfra/JenRik/internal/tomlLoader/tomlChecker"
	"github.com/Yohannfra/JenRik/internal/utils"
	"github.com/pelletier/go-toml"
	"log"
)

func LoadTestFile(fp string) tester.TestSuiteData {
	var testSuiteData tester.TestSuiteData
	var err error

	fc := utils.GetFileContent(fp)
	TomlContent, err := toml.Load(fc)
	testSuiteData.BinaryPath = ""

	if err != nil {
		log.Fatal(err)
	}
	for _, key := range TomlContent.Keys() {
		if key == "binary_path" {
			testSuiteData.BinaryPath = TomlContent.Get(key).(string)
		} else if key == "build_command" {
			err := utils.RunShellCommand(TomlContent.Get(key).(string))
			if err != nil {
				fmt.Println("build command failed: ", err)
			}
		} else {
			tomlChecker.CheckTestsValidity(key, TomlContent.Get(key).(*toml.Tree))
		}
	}
	if testSuiteData.BinaryPath == "" {
		log.Fatal("Could not find binary_path key in", fp)
	}

	testSuiteData.BinaryPath = tomlChecker.FixBinaryPath(testSuiteData.BinaryPath, fp)

	for _, key := range TomlContent.Keys() {
		if key == "binary_path" || key == "build_command" {
			continue
		}
		t := testData.NewTest(key, TomlContent.Get(key).(*toml.Tree))
		testSuiteData.TestSuite = append(testSuiteData.TestSuite, t)
	}
	tomlChecker.CheckBinaryValidity(testSuiteData.BinaryPath)
	return testSuiteData
}
