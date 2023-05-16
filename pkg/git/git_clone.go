package git

import (
	"fmt"
	"os"
	"roller/pkg/files"
	"roller/pkg/interaction"
)

// Clone clones the provided gitUrl to a temporary directory, and returns the path to that temporary directory.
//
// The reference can be left empty to clone the default branch.
func Clone(gitUrl string, gitReference string) string {
	gitDir := files.CreateTmpDir()
	CloneToDir(gitUrl, gitReference, gitDir)
	return gitDir
}

// CloneToDir Clones the URL to the provided destPath (directory). The directory is created if it doesn't exist
//
// The reference can be left empty to clone the default branch.
func CloneToDir(gitUrl string, gitReference string, destPath string) {
	// Check destPath exists, otherwise create it
	_, err := os.Stat(destPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(destPath, os.ModePerm)
		interaction.HandleError(err, true)
	}

	// Clone the repo, launch it as an interactive process for auth
	fmt.Printf("Cloning git repo, destPath=%s, url=%s\n", destPath, gitUrl)

	args := []string{"clone"}
	if len(gitReference) > 0 {
		args = append(args, "-b", gitReference)
	}
	args = append(args, gitUrl, ".")

	err = interaction.LaunchInteractiveProcess(destPath, "git", args...)
	if err != nil {
		fmt.Println("Failed to clone repository, aborted!")
		interaction.HandleError(err, false)
	}
}
