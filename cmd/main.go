package main

import (
	"fmt"
	"os"
	"roller/pkg/command"
	"roller/pkg/files"
)

func main() {
	var cmd = ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	defer func() {
		files.TmpDirCleanup()
		if err := recover(); err != nil {
			fmt.Printf("Error occurred: %s\n", err)
			os.Exit(1)
		}
	}()

	var success = false
	switch cmd {
	case "create":
		success = command.CommandCreate()
	case "update":
		success = command.CommandUpdate()
	case "version":
		success = command.CommandVersion()
	case "help":
		success = command.CommandHelp()
	default:
		var handled = command.CommandAction()
		if !handled {
			command.CommandHelp()
		}
	}

	if !success {
		os.Exit(1)
	}
}
