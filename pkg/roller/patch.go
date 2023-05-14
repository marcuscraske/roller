package roller

import (
	"fmt"
	"os"
	"path/filepath"
	"roller/pkg/git"
	"roller/pkg/interaction"
)

func Patch(firstTime bool, config Config, gitDir string, targetDir string) {

	// Copy tracked files from git clone to newChangesTmpDir
	newChangesTmpDir := CreateTmpDirAndCopyTrackedFiles(config, gitDir, gitDir)

	// Apply templating to newChangesTmpDir
	TemplatingApply(newChangesTmpDir, config)

	// Copy tracked files from target dir to oldChangesTmpDir
	oldChangesTmpDirPath := CreateTmpDirAndCopyTrackedFiles(config, targetDir, newChangesTmpDir)
	fmt.Println(oldChangesTmpDirPath)

	// Copy roller file for first time only
	if firstTime {
		CopyFile(gitDir+"/roller.yaml", newChangesTmpDir+"/roller.yaml")
	}

	// Generate diff
	diff := git.Diff(oldChangesTmpDirPath, newChangesTmpDir)

	// Check for deleted files, append to diff
	// TODO...

	// Display the changes
	fmt.Println("Changes:")
	fmt.Println("***************************************************************")
	fmt.Printf("%s\n\n", diff)
	fmt.Println("***************************************************************")
	fmt.Println("Apply? (y = apply, d = dump to 'patch.txt', n or anything else = abort)")

	var answer = interaction.PromptOrExit("y", "d")

	switch answer {
	case "y":
		// Apply patch
		git.Patch(diff)
	case "d":
		// Dump patch to target dir
		var data = []byte(diff)
		var err = os.WriteFile(targetDir+"/patch.txt", data, 0664)
		interaction.HandleError(err)
	}

	// Update list of tracked files
	var trackedFiles []string
	err := filepath.Walk(newChangesTmpDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			var relativePath = path[len(newChangesTmpDir)+1:]
			trackedFiles = append(trackedFiles, relativePath)
		}
		return nil
	})
	interaction.HandleError(err)

	// Write tracked files to target dir's roller.yaml
	UpdateTrackedFiles(targetDir, trackedFiles, config)
}
