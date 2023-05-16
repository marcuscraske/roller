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
	_, err = git.Status(targetDir)
	if err != nil {
		// Clone to a new folder, like a normal git clone
		folderName := getRepoName(gitUrl)
		gitDir := targetDir + "/" + folderName
		git.CloneToDir(gitUrl, gitReference, gitDir)
	} else {
		// Clone into the current directory and perform a patch
		gitDir := git.Clone(gitUrl, gitReference)
		config, err := ReadConfig(gitDir)
		interaction.HandleError(err, true)

		// Do the magic!
		Patch(config, gitDir, targetDir)
	}

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
