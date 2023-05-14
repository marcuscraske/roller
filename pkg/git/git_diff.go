package git

import (
	"fmt"
	"os/exec"
	"roller/pkg/interaction"
	"strings"
)

func Diff(srcDir string, destDir string) string {

	var args []string

	args = append(args, "-c", "diff.noprefix=true",
		"diff",
		"--no-index",
		"--relative",
		"--binary",
		"--no-ext-diff",
		"--no-color",
		"--src-prefix=old/",
		"--dst-prefix=new/",
		srcDir,
		destDir)

	process := exec.Command("git", args...)

	output, err := process.Output()

	// Git diff will give 0 for no changes and 1 for changes; thus ignore exit codes 0 and 1
	if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() > 1 {
		fmt.Printf("Running git diff failed, exitCode=%d\n", exitError.ExitCode())
		interaction.HandleError(err, false)
	}

	// Get the string output
	var result = string(output)

	// Check if there's no changes
	var trimmed = strings.TrimSpace(result)
	if len(trimmed) == 0 {
		return ""
	}

	// Replace the paths in the diff, as they'll have the full paths rather than a/ and b/
	result = strings.ReplaceAll(result, "old"+srcDir, "a")
	result = strings.ReplaceAll(result, "new"+destDir, "b")
	result = strings.ReplaceAll(result, "old"+destDir, "b")

	return result
}
