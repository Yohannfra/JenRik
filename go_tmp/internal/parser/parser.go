package parser

import (
	"github.com/Yohannfra/JenRik/internal/utils" // is_in
	"github.com/pelletier/go-toml"
	"log"
)

var TestsKeys = []string {
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

func CheckTestsValidity(testName string, testDict *toml.Tree) {
	requieredKeys := []string{"status", "args"}
	incompatiblesKeys := [][]string{
		{"stdout", "stdout_file"},
		{"stderr", "stderr_file"},
		{"stdin", "stdin_file"}}

	for _, key := range requieredKeys {
		if ! testDict.Has(key) {
			log.Fatalf("%s : Missing field : %s", testName, key)
		}
	}
	for _, key := range testDict.Keys() {
		if ! utils.IsIn(key, TestsKeys) {
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

