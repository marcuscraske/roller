package git

import (
	"fmt"
	"os/exec"
	"roller/pkg/interaction"
	"strings"
)

func Patch(workingDirectory string, patch string) {
	var args []string
	args = append(args,
		"apply", "-3",
	)

	process := exec.Command("git", args...)
	process.Dir = workingDirectory

	// Prepare stdin later for pushing the patch data
	process.Stdin = strings.NewReader(patch)

	// Capture the output
	output, err := process.Output()
	var result = string(output)
	fmt.Println(result)

	// Check patch applied successfully
	if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() > 0 {
		fmt.Printf("Running git apply failed, try dumping and manually applying the patch - exitCode=%d\n", exitError.ExitCode())
		interaction.HandleError(err, false)
	}

	fmt.Println("Patch applied!")
}
