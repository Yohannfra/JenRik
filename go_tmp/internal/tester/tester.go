package tester

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

const (
	ANSI_RED   = "\033[1;31m%s\033[0m"
	ANSI_GREEN = "\033[1;32m%s\033[0m"
	ANSI_RESET = "\033[1;39m%s\033[0m"
)

type TestSuiteData struct {
	BinaryPath  string
	TomlContent *toml.Tree
	TotalTests  int
	FailedTests int
}

func printSummary(data TestSuiteData) {
	fmt.Printf("\nSummary %s: %d tests ran\n", data.BinaryPath, data.TotalTests)
	fmt.Printf("%d : "+ANSI_GREEN+"\n", data.TotalTests-data.FailedTests, "OK")
	fmt.Printf("%d : "+ANSI_RED+"\n", data.FailedTests, "KO")
}

func Run(data TestSuiteData) {
	for _, key := range data.TomlContent.Keys() {
		if key == "binary_path" || key == "build_command" {
			continue

		}
		fmt.Println(key)
		data.TotalTests += 1
	}
	printSummary(data)
	os.Exit(data.FailedTests)
}
