package roller

import (
	"errors"
	"fmt"
	"os"
	"roller/pkg/git"
	"roller/pkg/interaction"
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
	var config, err2 = ReadConfig(gitDir)
	interaction.HandleError(err2)

	// Do the magic!
	Patch(true, config, gitDir, targetDir)
}
