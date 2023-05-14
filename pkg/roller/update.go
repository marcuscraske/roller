package roller

import (
	"errors"
	"fmt"
	"os"
	"roller/pkg/git"
	"roller/pkg/interaction"
)

func Update() {
	// Get working directory as target directory
	var targetDir, err = os.Getwd()
	interaction.HandleError(err)

	// Check roller config exists
	_, err = os.Stat(targetDir + "/" + ConfigFileName)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("No " + ConfigFileName + " in the current directory, unable to update - aborted!")
		return
	}

	// Clone the repo to tmp dir, read repo from roller config
	var config, err2 = ReadConfig(targetDir)
	interaction.HandleError(err2)

	var gitUrl = config.Template.Repo
	var gitDir = git.Clone(gitUrl)

	// Read the new config and merge it
	var newConfig, err3 = ReadConfig(gitDir)
	interaction.HandleError(err3)

	var mergedConfig = MergeConfig(newConfig, config)

	// Do the magic!
	Patch(mergedConfig, gitDir, targetDir)
}
