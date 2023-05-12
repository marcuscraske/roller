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
	var config = ReadConfig(targetDir)
	var gitUrl = config.Template.Repo
	var gitDir = git.Clone(gitUrl)

	// Read the new config and merge it
	var newConfig = ReadConfig(gitDir)
	config = MergeConfig(newConfig, config)

	// Do the magic!
	Patch(false, config, gitDir, targetDir)
}
