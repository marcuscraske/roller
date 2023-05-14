package git

import (
	"fmt"
	"os/exec"
	"roller/pkg/interaction"
	"strings"
)

func Patch(patch string) {
	var args []string
	args = append(args,
		"apply", "-3",
	)

	process := exec.Command("git", args...)

	// Prepare stdin later for pushing the patch data
	process.Stdin = strings.NewReader(patch)

	// Capture the output
	output, err := process.Output()
	var result = string(output)
	fmt.Println(result)

	// Git diff will give 0 for no changes and 1 for changes; thus ignore exit codes 0 and 1
	if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() > 1 {
		interaction.HandleError(err)
	}

	fmt.Println("Patch applied!")
}