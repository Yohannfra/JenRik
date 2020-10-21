package parser

import (
	"log"
	"github.com/pelletier/go-toml"
	"github.com/Yohannfra/Jenrik/internal/utils" // is_in
)

var TESTS_KEYS = []string {
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

func check_tests_validity(test_name string, test_dict *toml.Tree) {
	requiered_keys := []string{"status", "args"}
	incompatibles_keys := [][]string{
		{"stdout", "stdout_file"},
		{"stderr", "stderr_file"},
		{"stdin", "stdin_file"}}

	for _, key := range requiered_keys {
		if ! test_dict.Has(key) {
			log.Fatalf("%s : Missing field : %s", test_name, key)
		}
	}
	for _, key := range test_dict.Keys() {
		if ! is_in(key, TESTS_KEYS) {
			log.Fatalf("Unknown key: %s", key)
		}
		// TODO : Check type

		for _, ick := range incompatibles_keys {
			if key == ick[0] && is_in(ick[1], test_dict.Keys()) ||
			key == ick[1] && is_in(ick[0], test_dict.Keys()) {
				log.Fatalf("%s: Incompatible keys, %s and %s", test_name, ick[0], ick[1])
			}
		}
	}
}

