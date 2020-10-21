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

	if utils.Is_in("-q", argv) || utils.Is_in("--quiet", argv) {
		fmt.Println("Quiet mode")
		// TODO
	} else if argc == 2 && utils.Is_in("--version", argv) {
		fmt.Println(version.JENRIK_VERSION)
		os.Exit(0)
	} else if argc == 2 && utils.Is_in("--help", argv) {
		help.Print_help(argv[0])
		os.Exit(0)
	} else if argc == 3 && argv[1] == "init" {
		fmt.Println("Init")
	} else if argc == 2 && argv[1] != "init" {
		jenrik.Start(argv[1])
	} else {
		help.Print_help(argv[0])
		os.Exit(1)
	}
}
