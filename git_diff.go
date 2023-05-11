package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func GitDiff(srcDir string, destDir string, display bool) string {

	var args []string

	args = append(args, "-c", "diff.noprefix=true",
		"diff",
		"--no-index", "--relative", "--binary",
		"--no-ext-diff",
		"--src-prefix=old/",
		"--dst-prefix=new/",
		srcDir,
		destDir)

	if !display {
		args = append(args,
			"--no-color")
	}

	process := exec.Command("git", args...)

	fmt.Printf("%s, %s\n", srcDir, destDir)

	output, err := process.Output()

	// Git diff will give 0 for no changes and 1 for changes; thus ignore exit codes 0 and 1
	if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() > 1 {
		HandleError(err)
	}

	// Get the string output
	var result = string(output)

	// Replace the paths in the diff, as they'll have the full paths rather than a/ and b/
	result = strings.ReplaceAll(result, "old"+srcDir, "a")
	result = strings.ReplaceAll(result, "new"+destDir, "b")
	result = strings.ReplaceAll(result, "old"+destDir, "b")

	return result
}
