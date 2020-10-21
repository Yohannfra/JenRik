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

func CheckTestsValidity(testName string, testDict *toml.Tree) {
	requiredKeys := []string{"status", "args"}
	incompatiblesKeys := [][]string{
		{"stdout", "stdout_file"},
		{"stderr", "stderr_file"},
		{"stdin", "stdin_file"}}

	TestsKeys := []string{
		"args",
		"status",
		"stdout",
		"stderr",
		"pre",
		"post",
		"stdout_file",
		"stderr_file",
		"pipe_stdout",
		"pipe_stderr",
		"timeout",
		"should_fail",
		"stdin",
		"stdin_file",
		"env",
		"add_env",
		"repeat",
	}

	for _, key := range requiredKeys {
		if !testDict.Has(key) {
			log.Fatalf("%s : Missing field : %s", testName, key)
		}
	}
	for _, key := range testDict.Keys() {
		if !utils.IsIn(key, TestsKeys) {
			log.Fatalf("Unknown key: %s", key)
		}
		// TODO : Check type

		for _, ick := range incompatiblesKeys {
			if key == ick[0] && utils.IsIn(ick[1], testDict.Keys()) ||
				key == ick[1] && utils.IsIn(ick[0], testDict.Keys()) {
				log.Fatalf("%s: Incompatible keys, %s and %s", testName, ick[0], ick[1])
			}
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

	fmt.Printf("Running build command : '%s'\n", command)
	err := cmd.Run()
	if err != nil {
		log.Println("Error running build command: ", err)
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
