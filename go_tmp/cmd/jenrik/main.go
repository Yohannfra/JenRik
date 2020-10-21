package main

import (
	"fmt"
	"github.com/Yohannfra/JenRik/internal/help"
	"github.com/Yohannfra/JenRik/internal/jenrik"
	"github.com/Yohannfra/JenRik/internal/utils"
	"github.com/Yohannfra/JenRik/internal/version"
	"log"
	"os"
)

func main() {
	argv := os.Args
	argc := len(argv)
	log.SetFlags(0)

	if utils.IsIn("-q", argv) || utils.IsIn("--quiet", argv) {
		fmt.Println("Quiet mode")
		// TODO
	} else if argc == 2 && utils.IsIn("--version", argv) {
		fmt.Println(version.JenrikVersion)
		os.Exit(0)
	} else if argc == 2 && utils.IsIn("--help", argv) {
		help.PrintHelp(argv[0])
		os.Exit(0)
	} else if argc == 3 && argv[1] == "init" {
		fmt.Println("Init")
	} else if argc == 2 && argv[1] != "init" {
		jenrik.Start(argv[1])
	} else {
		help.PrintHelp(argv[0])
		os.Exit(1)
	}
}
