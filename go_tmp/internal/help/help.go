package help

import "fmt"

func print_help(binary_name string) {
    // Print a basic help showing how to use Jenerik
    fmt.Printf("USAGE : %s file.jrk | init path_to_binary\n", binary_name)
    fmt.Println("\tinit\t\tcreate a basic test file for the given binary")
    fmt.Println("\t--version\tprint version information and exit")
    fmt.Println("\t--help\tprint this help and exit")
    fmt.Println("\t--quiet, -q\trun in quiet mode (doesn't show the diffs)")
}
