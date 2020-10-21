package main

import (
	"fmt"
	"os"
	"log"
	"github.com/Yohannfra/JenRik/internal/utils"
	"github.com/Yohannfra/JenRik/internal/version"
	"github.com/Yohannfra/JenRik/internal/help"
	"github.com/Yohannfra/JenRik/internal/jenrik"
)

func main() {
	argv := os.Args
	argc := len(argv)
	log.SetFlags(0)

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
	} else if argc == 2 && argv[1] != "init" {
		start_jenrik(argv[1])
	} else {
		print_help(argv[0])
		os.Exit(1)
	}
}
