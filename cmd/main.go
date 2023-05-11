package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"roller/pkg/command"
	"roller/pkg/interaction"
	"roller/pkg/roller"
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

func LaunchInteractiveProcess(name string, arg ...string) {
	fmt.Println("Launching vim for roller.yaml...")
	process := exec.Command(name, arg...)
	stdin, err := process.StdinPipe()
	interaction.HandleError(err)

	defer stdin.Close()
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	if err = process.Start(); err != nil {
		interaction.HandleError(err)
	}

	err = process.Wait()
	if err != nil {
		interaction.HandleError(err)
	}
}

// ApplyTemplating Applies templating to the provided dir
func ApplyTemplating(dir string) {
	// Read config for vars
	config := roller.ReadConfig(dir + "/roller.yaml")

	// Setup context
	//context := pongo2.Context{"vars": config.template.vars}

	// Iterate and apply template to each file
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		return ApplyTemplatingFile(dir, path, info, config)
	})
	interaction.HandleError(err)
}

func ApplyTemplatingFile(dir string, path string, info os.FileInfo, config roller.Config) error {
	//if IsIgnoredFile(config, dir, path, info) {
	//	fmt.Printf("Skipped templating, file=%s\n", path)
	//	return nil
	//}

	fmt.Printf("Processing, file=%s\n", path)

	// Apply templating

	// Check whether new/diff to target file

	return nil
}
