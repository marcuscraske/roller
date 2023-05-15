package git

import (
	"fmt"
	"os"
	"os/exec"
	"roller/pkg/interaction"
	"strings"
)

func Status(dir string) string {
	// Check git exists, otherwise return empty string
	_, err := os.Stat(dir + "/.git")
	if os.IsNotExist(err) {
		fmt.Println("The current working directory must be a git repository, run git init if this is a new directory!")
		interaction.HandleError(err, false)
		return "no git folder"
	}

	// Run git status to check for the current changes
	process := exec.Command("git", "status", "--porcelain")
	process.Dir = dir
	output, err := process.Output()
	if err != nil {
		fmt.Println("Running git status failed")
		interaction.HandleError(err, false)
	}

	// Get the string output
	var result = string(output)

	// Check if there's no changes
	var trimmed = strings.TrimSpace(result)
	if len(trimmed) == 0 {
		return ""
	}
	return result
}
