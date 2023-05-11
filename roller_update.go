package main

import (
	"errors"
	"fmt"
	"os"
)

func RollerUpdate(firstTime bool) {
	// Get working directory as target directory
	var targetDir, err = os.Getwd()
	HandleError(err)

	// Check roller config doesn't already exist in the current folder
	_, err = os.Stat(targetDir + "/roller.yaml")
	if !errors.Is(err, os.ErrNotExist) {
		fmt.Println("roller.yaml detected in current directory, aborted!")
		return
	}

	// Clone the repo to tmp dir / read roller config
	var gitDir string
	var config RollerConfig

	if firstTime {
		// TODO read repo url from first param
		gitDir = GitClone("https://github.com/marcuscraske/tmp.git")
		config = ReadConfig(gitDir + "/roller.yaml")
	} else {
		config := ReadConfig(targetDir + "/roller.yaml")
		gitUrl := config.template.repo
		gitDir = GitClone(gitUrl)
	}

	// Copy tracked files from git clone to tmpDir
	newChangesTmpDir := CreateTmpDirAndCopyTrackedFiles(config, gitDir, gitDir)

	// Apply templating to git clone files
	TemplatingApply(newChangesTmpDir)

	// Copy tracked files from target dir to old tmp dir
	oldChangesTmpDirPath := CreateTmpDirAndCopyTrackedFiles(config, targetDir, newChangesTmpDir)
	fmt.Println(oldChangesTmpDirPath)

	// Copy roller file for first time only
	if firstTime {
		CopyFile(gitDir+"/roller.yaml", newChangesTmpDir+"/roller.yaml")
	}

	// Generate diff
	diff := GitDiff(oldChangesTmpDirPath, newChangesTmpDir, true)

	fmt.Println("Changes:")
	fmt.Println("***************************************************************")
	fmt.Printf("%s\n\n", diff)
	fmt.Println("***************************************************************")
	fmt.Println("Apply? (y/n)")

	PromptYesOrExit()

	// Apply patch
	GitPatch(diff)

	// Update list of tracked files
	// TODO TBC...
}
