package main

import (
	"fmt"
	"github.com/Yohannfra/JenRik/internal/help"
	"github.com/Yohannfra/JenRik/internal/jenrik"
	"github.com/Yohannfra/JenRik/internal/logLevel"
	"github.com/Yohannfra/JenRik/internal/version"
	"log"
	"os"
)

func main() {
	argv := os.Args

	log.SetFlags(0) // disable time in logs

	for i, arg := range argv {
		if arg == "-q" || arg == "--quiet" {
			logLevel.LOG_LEVEL = logLevel.QUIET
			argv = append(argv[:i], argv[i+1:]...)
		} else if arg == "-d" || arg == "--debug" {
			logLevel.LOG_LEVEL = logLevel.DEBUG
			argv = append(argv[:i], argv[i+1:]...)
		} else if arg == "--version" {
			fmt.Println(version.JenrikVersion)
			os.Exit(0)
		} else if arg == "--help" {
			help.PrintHelp(argv[0])
			os.Exit(0)
		}
	}

	if len(argv) == 3 && argv[1] == "init" {
		jenrik.Init(argv[2])
	} else if len(argv) == 2 {
		jenrik.Start(argv[1])
	} else {
		help.PrintHelp(argv[0])
		os.Exit(1)
	}
}
