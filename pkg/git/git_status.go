package git

import (
	"os/exec"
	"roller/pkg/interaction"
	"strings"
)

func Status(dir string) string {
	// Run git status to check for the current changes
	process := exec.Command("git", "status", "--porcelain")
	output, err := process.Output()
	interaction.HandleError(err)

	// Get the string output
	var result = string(output)

	// Check if there's no changes
	var trimmed = strings.TrimSpace(result)
	if len(trimmed) == 0 {
		return ""
	}
	return result
}
