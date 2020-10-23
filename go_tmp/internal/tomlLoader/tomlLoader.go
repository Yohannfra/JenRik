package tomlLoader

import (
	"github.com/Yohannfra/JenRik/internal/tester"
	"github.com/Yohannfra/JenRik/internal/tomlLoader/tomlChecker"
	"github.com/Yohannfra/JenRik/internal/utils"
	"github.com/pelletier/go-toml"
	"log"
)

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
			tomlChecker.RunBuildCommand(testData.TomlContent.Get(key).(string))
		} else {
			tomlChecker.CheckTestsValidity(key, testData.TomlContent.Get(key).(*toml.Tree))
		}
	}
	if testData.BinaryPath == "" {
		log.Fatal("Could not find binary_path key in", fp)
	}

	testData.BinaryPath = tomlChecker.FixBinaryPath(testData.BinaryPath, fp)
	tomlChecker.CheckBinaryValidity(testData.BinaryPath)
	return testData
}
