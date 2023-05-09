package main

import "fmt"

func CommandUpdate() {
	fmt.Println("Roller updating...")

	// Clone the git repo to tmp dir
	//tmpDir := GitClone("https://github.com/marcuscraske/tmp.git")

	// Apply templating to tmp dir
	//ApplyTemplating(tmpDir)

	// Copy tmpdir files (except those ignored), with templating applied

	// Compute git diff between the tmp and tracked dir

	// Prompt user for any changes requiring input

	// Copy changes

	// TODO launch interactive vim when template.vars changes
	//LaunchInteractiveProcess("vim", tmpDir+"/roller.yaml")

	fmt.Println("Finished!")
}
