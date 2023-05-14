package roller

import (
	"errors"
	"fmt"
	"os"
	"roller/pkg/git"
	"roller/pkg/interaction"
)

func Create(gitUrl string) bool {

	var targetDir, err = os.Getwd()

	// Check roller config doesn't already exist in the current folder
	_, err = os.Stat(targetDir + "/" + ConfigFileName)
	if !errors.Is(err, os.ErrNotExist) {
		fmt.Println("roller.yaml detected in current directory, aborted!")
		return false
	}

	// Perform the initial clone
	var gitDir = git.Clone(gitUrl)
	config, err := ReadConfig(gitDir)
	interaction.HandleError(err, true)

	// Do the magic!
	Patch(config, gitDir, targetDir)
	return true
}
