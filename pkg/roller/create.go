package roller

import (
	"errors"
	"fmt"
	"os"
	"roller/pkg/git"
)

func Create(gitUrl string) {

	var targetDir, err = os.Getwd()

	// Check roller config doesn't already exist in the current folder
	_, err = os.Stat(targetDir + "/roller.yaml")
	if !errors.Is(err, os.ErrNotExist) {
		fmt.Println("roller.yaml detected in current directory, aborted!")
		return
	}

	// Perform the initial clone
	var gitDir = git.Clone(gitUrl)
	var config = ReadConfig(gitDir)

	// Do the magic!
	Patch(true, config, gitDir, targetDir)
}
