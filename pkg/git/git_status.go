package git

import (
	"fmt"
	"os"
	"os/exec"
	"roller/pkg/interaction"
	"strings"
)

func Status(dir string) (result string, err error) {
	// Check git exists, otherwise return empty string
	_, err = os.Stat(dir + "/.git")
	if os.IsNotExist(err) {
		return "", err
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
	result = string(output)

	// Check if there's no changes
	var trimmed = strings.TrimSpace(result)
	if len(trimmed) == 0 {
		return "", nil
	}
	return result, nil
}
