package main

import (
	"fmt"
	"os"
)

// import "github.com/pelletier/go-toml"

const JENRIK_VERSION = "1.10"

func print_help(binary_name string) {
    // Print a basic help showing how to use Jenerik
    fmt.Printf("USAGE : %s file.jrk | init path_to_binary\n", binary_name)
    fmt.Println("\tinit\t\tcreate a basic test file for the given binary")
    fmt.Println("\t--version\tprint version information and exit")
    fmt.Println("\t--help\tprint this help and exit")
    fmt.Println("\t--quiet, -q\trun in quiet mode (doesn't show the diffs)")
}

func is_in(str string, arr []string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func main() {
	argv := os.Args
	argc := len(argv)

	if is_in("-q", argv) || is_in("--quiet", argv) {
		fmt.Println("Quiet mode")
		// TODO
	} else if argc == 2 && is_in("--version", argv) {
		fmt.Println(JENRIK_VERSION)
		os.Exit(0)
	} else if argc == 2 && is_in("--help", argv) {
		print_help(argv[0])
		os.Exit(0)
	} else if argc == 3 && argv[1] == "init" {
		fmt.Println("Init")
	} else if argc == 2 {
		fmt.Println("Start")
	} else {
		print_help(argv[0])
		os.Exit(1)
	}
}


//     n := map[string]int{"foo": 1, "bar": 2}
// 	fmt.Print(n)

// 	fmt.Println(argv)
// 	fmt.Println("CoucOu")
// 	config, _ := toml.Load(`
// 	[postgres]
// 	user = "pelletier"
// 	password = "mypassword"`)

// 	user := config.Get("postgres.user").(string)
// 	fmt.Print(user)
