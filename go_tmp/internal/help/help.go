package help

import "fmt"

func PrintHelp(binaryName string) {
	fmt.Printf("USAGE : %s file.jrk | init path_to_binary\n", binaryName)
	fmt.Println("\tinit\t\tcreate a basic test file for the given binary")
	fmt.Println("\t--version\tprint version information and exit")
	fmt.Println("\t--help\tprint this help and exit")
	fmt.Println("\t--quiet, -q\trun in quiet mode (doesn't show the diffs)")
	fmt.Println("\t--debug, -d\trun in debug mode")
}
