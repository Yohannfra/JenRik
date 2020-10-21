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
	quietMode := false
	log.SetFlags(0)

	if utils.IsIn("-q", argv) || utils.IsIn("--quiet", argv) {
		quietMode = true
	} else if argc == 2 && utils.IsIn("--version", argv) {
		fmt.Println(version.JenrikVersion)
		os.Exit(0)
	} else if argc == 2 && utils.IsIn("--help", argv) {
		help.PrintHelp(argv[0])
		os.Exit(0)
	} else if argc == 3 && argv[1] == "init" {
		jenrik.Init(argv[2])
	} else if argc == 2 && argv[1] != "init" {
		jenrik.Start(argv[1], quietMode)
	} else {
		help.PrintHelp(argv[0])
		os.Exit(1)
	}
}
