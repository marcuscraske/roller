package roller

import (
	"errors"
	"fmt"
	"os"
	"roller/pkg/git"
	"roller/pkg/interaction"
	"strings"
)

func Create(gitUrl string, gitReference string) bool {

	var targetDir, err = os.Getwd()

	// Check roller config doesn't already exist in the current folder
	_, err = os.Stat(targetDir + "/" + ConfigFileName)
	if !errors.Is(err, os.ErrNotExist) {
		fmt.Println("roller.yaml detected in current directory, aborted!")
		return false
	}

	// Check whether git exists already, otherwise clone to new directory and skip patch process
	var gitDir string

	_, err = git.Status(targetDir)
	if err != nil {
		// Go to clone as if it's git clone (new sub-directory)
		folderName := getRepoName(gitUrl)
		gitDir = targetDir + "/" + folderName
		_, err = git.Status(gitDir)
		if err != nil {
			git.CloneToDir(gitUrl, gitReference, gitDir)
			targetDir = gitDir
		} else {
			fmt.Println("Unable to clone template, conflicting sub-directory exists at '" + folderName + "'")
			fmt.Println("")
			fmt.Println("Please cd into the directory '" + folderName + "', and run this again to apply to an existing git repo.")
			return false
		}
	} else {
		// Clone into the current directory
		gitDir = git.Clone(gitUrl, gitReference)
	}

	// Read config
	config, err := ReadConfig(gitDir)
	interaction.HandleError(err, true)

	// Perform a survey
	if !Survey(&config, nil, true) {
		return false
	}

	// Do the magic!
	Patch(config, gitDir, targetDir)
	return true
}

func getRepoName(gitUrl string) (result string) {
	result = "repo"
	i := strings.LastIndex(gitUrl, "/")
	if i > 0 && i < len(gitUrl) {
		result = gitUrl[i+1:]
		if strings.HasSuffix(result, ".git") {
			result = result[:len(result)-4]
		}
	}
	return result
}
