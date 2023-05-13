package roller

import (
	"os"
	"roller/pkg/git"
	"roller/pkg/interaction"
)

func Update() {
	// Get working directory as target directory
	var targetDir, err = os.Getwd()
	interaction.HandleError(err)

	// Clone the repo to tmp dir, read repo from roller config
	var config, err2 = ReadConfig(targetDir)
	interaction.HandleError(err2)

	var gitUrl = config.Template.Repo
	var gitDir = git.Clone(gitUrl)

	// Read the new config and merge it
	var newConfig, err3 = ReadConfig(gitDir)
	interaction.HandleError(err3)

	config = MergeConfig(newConfig, config)

	// Do the magic!
	Patch(false, config, gitDir, targetDir)
}
