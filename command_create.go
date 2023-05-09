package main

import (
	"errors"
	"fmt"
	"os"
)

func CommandCreate() {

	// Get working directory as target directory
	var targetDir, err = os.Getwd()
	HandleError(err)

	// Check roller config doesn't already exist in the current folder
	_, err = os.Stat(targetDir + "/roller.yaml")
	if !errors.Is(err, os.ErrNotExist) {
		fmt.Println("roller.yaml detected in current directory, aborted!")
		return
	}

	// Clone the repo to tmp dir
	gitDir := GitClone("https://github.com/marcuscraske/tmp.git")

	// Read the roller config
	config := ReadConfig(gitDir + "/roller.yaml")

	// Copy tracked files from git clone to tmpDir
	newChangesTmpDir := CreateTmpDirAndCopyTrackedFiles(config, gitDir, gitDir)

	// Apply templating to pulled files
	TemplatingApply(newChangesTmpDir)

	// Copy tracked files from target dir to old tmp dir
	oldChangesTmpDirPath := CreateTmpDirAndCopyTrackedFiles(config, targetDir, newChangesTmpDir)
	fmt.Println(oldChangesTmpDirPath)

	// Copy roller file, as it's the first time
	CopyFile(gitDir+"/roller.yaml", newChangesTmpDir+"/roller.yaml")

	// Generate diff
	// -- We do it twice as one is computer
	diff := GitDiff(oldChangesTmpDirPath, newChangesTmpDir, false)
	fmt.Println("Differences:")
	fmt.Printf("%s\n\n", diff)
	fmt.Println("Apply? (y/n)")

}
