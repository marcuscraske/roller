package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func Reference(dir string) (result string, err error) {
	process := exec.Command("git", "rev-parse", "HEAD")
	process.Dir = dir
	output, err := process.Output()
	if err != nil {
		fmt.Println("Running git reference failed")
		return "", err
	}

	result = string(output)
	result = strings.TrimSpace(result)
	return result, nil
}
