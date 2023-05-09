package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var command = ""
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "create":
		CommandCreate()
	case "update":
		CommandUpdate()
	case "version":
		CommandVersion()
	default:
		CommandHelp()
	}
}

func LaunchInteractiveProcess(name string, arg ...string) {
	fmt.Println("Launching vim for roller.yaml...")
	process := exec.Command(name, arg...)
	stdin, err := process.StdinPipe()
	HandleError(err)

	defer stdin.Close()
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	if err = process.Start(); err != nil {
		HandleError(err)
	}

	err = process.Wait()
	if err != nil {
		HandleError(err)
	}
}

// ApplyTemplating Applies templating to the provided dir
func ApplyTemplating(dir string) {
	// Read config for vars
	config := ReadConfig(dir + "/roller.yaml")

	// Setup context
	//context := pongo2.Context{"vars": config.template.vars}

	// Iterate and apply template to each file
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		return ApplyTemplatingFile(dir, path, info, config)
	})
	HandleError(err)
}

func ApplyTemplatingFile(dir string, path string, info os.FileInfo, config RollerConfig) error {
	//if IsIgnoredFile(config, dir, path, info) {
	//	fmt.Printf("Skipped templating, file=%s\n", path)
	//	return nil
	//}

	fmt.Printf("Processing, file=%s\n", path)

	// Apply templating

	// Check whether new/diff to target file

	return nil
}

// HandleError Checks whether the error is not nil; if so, it's printed to stdout and the program exits
func HandleError(err error) {
	// TODO use panic(err) instead?
	if err != nil {
		fmt.Printf("Failed due to unexpected error: %s", err)
		os.Exit(1)
	}
}
