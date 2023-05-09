package main

import (
	"fmt"
	"os/exec"
)

func GitDiff(srcDir string, destDir string, color bool) string {

	var args []string

	args = append(args, "-c", "diff.noprefix=true",
		"diff",
		"--no-index", "--relative", "--binary",
		"--no-ext-diff",
		"--src-prefix="+srcDir,
		"--dst-prefix="+destDir,
		srcDir,
		destDir)

	if !color {
		args = append(args, "--no-color")
	}

	process := exec.Command("git", args...)

	fmt.Printf("%s, %s\n", srcDir, destDir)

	output, err := process.Output()

	// Git diff will give 0 for no changes and 1 for changes; thus ignore exit codes 0 and 1
	if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() > 1 {
		HandleError(err)
	}

	var result = string(output)
	return result
}
