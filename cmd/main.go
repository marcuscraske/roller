package main

import (
	"os"
	"roller/pkg/command"
)

func main() {
	var cmd = ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "create":
		command.CommandCreate()
	case "update":
		command.CommandUpdate()
	case "version":
		command.CommandVersion()
	default:
		command.CommandHelp()
	}
}
