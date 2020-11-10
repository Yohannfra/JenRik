package tester

import (
	"bytes"
	"fmt"
	"github.com/Yohannfra/JenRik/internal/testData"
	"github.com/Yohannfra/JenRik/internal/utils"
	"log"
	"os"
	"time"
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
	TestSuite   []*testData.Test
	TotalTests  int
	FailedTests int
}

type ShellCommandData struct {
	exitStatus int
	stdout     string
	stderr     string
	timeout    bool
}

func printSummary(testSuiteData *TestSuiteData) {
	fmt.Printf("\nSummary %s: %d tests ran\n", testSuiteData.BinaryPath, testSuiteData.TotalTests)
	fmt.Printf("%d : "+ANSI_GREEN+"\n", testSuiteData.TotalTests-testSuiteData.FailedTests, "OK")
	fmt.Printf("%d : "+ANSI_RED+"\n", testSuiteData.FailedTests, "KO")
}

func runWithTimeout(cmd *exec.Cmd, timeout_time int) (bool, error) {
	done := make(chan error)

	go func() {
		done <- cmd.Run()
	}()

	timeout := time.After(1 * time.Second)

	select {
	case <-timeout:
		err := cmd.Process.Kill()
		if err != nil {
			log.Fatal("Error, could not kill timout process")
		}
		return true, nil
	case err := <-done:
		return false, err
	}
}

func runCmd(command string, env map[string]string, timeout_time int) ShellCommandData {
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

	for key, value := range env {
		cmd.Env = append(cmd.Env, key+"="+value)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if timeout_time > 0 {
		res, err := runWithTimeout(cmd, timeout_time)
		data.timeout = res
		if err != nil {
			data.exitStatus = err.(*exec.ExitError).ExitCode()
		}
	} else {
		data.timeout = false
		err := cmd.Run()
		if err != nil {
			data.exitStatus = err.(*exec.ExitError).ExitCode()
		}
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

func checkTestResult(test *testData.Test, testRes *ShellCommandData) bool {
	// timeout
	if testRes.timeout == true {
		return printTestFail("Test timed out after  %d seconds\n", test.Timeout)
	}

	// exit status
	if test.Status != testRes.exitStatus {
		return printTestFail("Invalid exit status, expected %d but got %d\n", test.Status, testRes.exitStatus)
	}

	// stdout
	if test.Stdout != "" {
		if test.Stdout != testRes.stdout {
			printTestFail("Invalid stdout\n")
			utils.PrintDiff(test.Stdout, testRes.stdout)
			return false
		}
	}
	// stderr
	if test.Stderr != "" {
		if test.Stderr != testRes.stderr {
			printTestFail("Invalid stderr\n")
			utils.PrintDiff(test.Stderr, testRes.stderr)
			return false
		}
	}
	return true
}

func runTest(binaryPath string, test *testData.Test) bool {
	args := test.Args

	if test.Pre != "" {
		err := utils.RunShellCommand(test.Pre)
		if err != nil {
			fmt.Println("Pre command error: ", err)
		}
	}
	testResult := runCmd(binaryPath+" "+strings.Join(args, " "), test.Env, test.Timeout)
	if test.Post != "" {
		err := utils.RunShellCommand(test.Post)
		if err != nil {
			fmt.Println("Post command error: ", err)
		}
	}
	res := checkTestResult(test, &testResult)

	// should fail flag
	if test.ShouldFail && res {
		return false
	} else if test.ShouldFail && !res {
		return false
	}
	return res
}

func repeatTest(binaryPath string, test *testData.Test) int {
	failedTests := 0
	testToRun := test.Repeat

	fmt.Print("\n")
	for test.Repeat > 0 {
		fmt.Printf("  -> repeat %d : ", testToRun-test.Repeat+1)
		if !runTest(binaryPath, test) {
			failedTests += 1
		} else {
			fmt.Printf(ANSI_GREEN+"\n", "OK")
		}
		test.Repeat -= 1
	}
	return failedTests
}

func Run(testSuiteData *TestSuiteData) {
	for _, test := range testSuiteData.TestSuite {
		fmt.Print(test.Name, ": ")
		if test.Repeat > 0 {
			testSuiteData.TotalTests += test.Repeat
			testSuiteData.FailedTests += repeatTest(testSuiteData.BinaryPath, test)
			continue
		}
		if !runTest(testSuiteData.BinaryPath, test) {
			testSuiteData.FailedTests += 1
		} else {
			fmt.Printf(ANSI_GREEN+"\n", "OK")
		}
		testSuiteData.TotalTests += 1
	}
	printSummary(testSuiteData)
	os.Exit(testSuiteData.FailedTests)
}
