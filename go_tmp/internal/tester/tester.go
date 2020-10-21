package tester

import "github.com/pelletier/go-toml"

type TestSuiteData struct {
	BinaryPath  string
	TomlContent *toml.Tree
	TotalTests  uint
	FailedTests uint
}

func Run(data TestSuiteData) {

}
