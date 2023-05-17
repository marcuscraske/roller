package roller

import (
	"errors"
	"fmt"
	"os"
	"roller/pkg/git"
	"roller/pkg/interaction"
)

func Update(gitReference string) bool {
	// Get working directory as target directory
	targetDir, err := os.Getwd()
	interaction.HandleError(err, true)

	// Check roller config exists
	_, err = os.Stat(targetDir + "/" + ConfigFileName)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("No " + ConfigFileName + " in the current directory, unable to update - aborted!")
		return false
	}

	// Clone the repo to tmp dir, read repo from roller config
	config, err := ReadConfig(targetDir)
	interaction.HandleError(err, true)

	var gitUrl = config.Template.Repo
	var gitDir = git.Clone(gitUrl, gitReference)

	// Read the new config and merge it
	newConfig, err := ReadConfig(gitDir)
	interaction.HandleError(err, true)

	mergeResult := MergeConfig(newConfig, config)
	mergedConfig := mergeResult.config

	// Perform a survey (abort if survey fails)
	if !Survey(&mergedConfig, mergeResult.changedVars, false) {
		return false
	}

	// Do the magic!
	return Patch(mergedConfig, gitDir, targetDir)
}
