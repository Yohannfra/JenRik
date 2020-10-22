package tester

import (
	"bytes"
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
	"os/exec"
	"strings"
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

type ShellCommandData struct {
	exitStatus int
	stdout     string
	stderr     string
}

func printSummary(testSuiteData TestSuiteData) {
	fmt.Printf("\nSummary %s: %d tests ran\n", testSuiteData.BinaryPath, testSuiteData.TotalTests)
	fmt.Printf("%d : "+ANSI_GREEN+"\n", testSuiteData.TotalTests-testSuiteData.FailedTests, "OK")
	fmt.Printf("%d : "+ANSI_RED+"\n", testSuiteData.FailedTests, "KO")
}

func runCmd(command string) ShellCommandData {
	tmp := strings.Split(command, " ")
	var cmd *exec.Cmd
	var data ShellCommandData
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	if len(tmp) == 1 {
		cmd = exec.Command(tmp[0])
	} else {
		cmd = exec.Command(tmp[0], strings.Join(tmp[1:], " "))
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		data.exitStatus = err.(*exec.ExitError).ExitCode()
	}
	data.stderr = string(stderr.Bytes())
	data.stdout = string(stdout.Bytes())
	return data
}

func runTest(binaryPath string, testName string, testData *toml.Tree) bool {
	a := runCmd(binaryPath)
	fmt.Print(a)
	//argsTmp := testData.Get("args").([]interface{})
	//var args []string
	//
	//for _, arg := range argsTmp {
	//	argType := fmt.Sprintf("%T", arg)
	//	if argType == "string" {
	//		args = append(args, arg.(string))
	//	} else {
	//		log.Fatalf("Error, invalid type in args : %s", argType)
	//	}
	//}
	//fmt.Println(args)

	return false
}

func Run(testSuiteData TestSuiteData) {
	for _, key := range testSuiteData.TomlContent.Keys() {
		if key == "binary_path" || key == "build_command" {
			continue
		}
		fmt.Printf("%s : ", key)
		if !runTest(testSuiteData.BinaryPath, key, testSuiteData.TomlContent.Get(key).(*toml.Tree)) { // test fail
			testSuiteData.FailedTests += 1
			fmt.Printf(ANSI_RED+"\n", "KO")
		} else { // test success
			fmt.Printf(ANSI_GREEN+"\n", "OK")
		}
		testSuiteData.TotalTests += 1
	}
	printSummary(testSuiteData)
	os.Exit(testSuiteData.FailedTests)
}
