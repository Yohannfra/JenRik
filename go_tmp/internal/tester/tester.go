package tester

import (
	"bytes"
	"fmt"
	"github.com/Yohannfra/JenRik/internal/logLevel"
	"github.com/pelletier/go-toml"
	"math"
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

func printSummary(testSuiteData *TestSuiteData) {
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

func printTestFail(format string, a ...interface{}) bool {
	fmt.Printf(ANSI_RED+"", "KO")
	fmt.Print(" : ")
	fmt.Printf(format, a...)
	return false
}

func printDiff(gotStr string, expectedStr string) {
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

func checkTestResult(testData *toml.Tree, testRes *ShellCommandData) bool {
	// exit status
	st, _ := testData.Get("status").(int64)
	if int64(testRes.exitStatus) != st {
		return printTestFail("Invalid exit status, expected %d but got %d\n", int(st), testRes.exitStatus)
	}

	// stdout
	if testData.Has("stdout") {
		val := testData.Get("stdout").(string)
		if val != testRes.stderr {
			printTestFail("Invalid stdout\n")
			printDiff(val, testRes.stderr)
			return false
		}
	}

	// stderr
	if testData.Has("stderr") {
		//if
		val := testData.Get("stderr").(string)
		if val != testRes.stderr {
			printTestFail("Invalid stderr\n")
			printDiff(val, testRes.stderr)
			return false
		}
	}
	return true
}

func runTest(binaryPath string, testName string, testData *toml.Tree) bool {
	var args []string

	argsTmp := testData.Get("args").([]interface{})
	for _, arg := range argsTmp {
		args = append(args, arg.(string))
	}
	testResult := runCmd(binaryPath + " " + strings.Join(args, " "))

	if testData.Has("repeat") {
		val := int(testData.Get("repeat").(int64))
		if val > 0 {
			fmt.Printf(" - Repeat %d %s: ", val, testName)

			testData.Set("repeat", int64(val-1))
			runTest(binaryPath, testName, testData)
		}
	}
	return checkTestResult(testData, &testResult)

}

func Run(testSuiteData *TestSuiteData) {
	for _, key := range testSuiteData.TomlContent.Keys() {
		if key == "binary_path" || key == "build_command" {
			continue
		}
		fmt.Printf("%s : ", key)
		if !runTest(testSuiteData.BinaryPath, key, testSuiteData.TomlContent.Get(key).(*toml.Tree)) { // test fail
			testSuiteData.FailedTests += 1
		} else { // test success
			fmt.Printf(ANSI_GREEN+"\n", "OK")
		}
		testSuiteData.TotalTests += 1
	}
	printSummary(testSuiteData)
	os.Exit(testSuiteData.FailedTests)
}
