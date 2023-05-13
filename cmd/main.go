package main

import (
	"os"
	"roller/pkg/command"
	"roller/pkg/files"
)

func main() {
	var cmd = ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	defer files.TmpDirCleanup()

	switch cmd {
	case "create":
		command.CommandCreate()
	case "update":
		command.CommandUpdate()
	case "version":
		command.CommandVersion()
	default:
		var handled = command.CommandAction()
		if !handled {
			command.CommandHelp()
		}
	}
}
