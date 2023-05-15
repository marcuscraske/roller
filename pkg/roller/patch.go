package roller

import (
	"fmt"
	"os"
	"path/filepath"
	"roller/pkg/files"
	"roller/pkg/git"
	"roller/pkg/interaction"
)

func Patch(config Config, gitDir string, targetDir string) bool {
	// Check targetDir has no pending git changes
	status := git.Status(targetDir)
	if len(status) > 0 {
		fmt.Println("Unable to proceed, git changes are pending - check git status!")
		return false
	}

	// Setup template context
	templateContext := CreateTemplateContext(config)

	// Setup directory for new changes
	newChangesTmpDir := files.CreateTmpDir()

	// Execute pre-actions
	ExecuteActions(newChangesTmpDir, config.Template.Actions.Pre, templateContext, config)

	// Copy tracked files from git clone to newChangesTmpDir
	MirrorTmpDirFilesFromTargetDir(config, gitDir, gitDir, newChangesTmpDir)

	// Apply templating to newChangesTmpDir
	ApplyTemplatingDir(newChangesTmpDir, config, templateContext)

	// Copy tracked files from target dir to oldChangesTmpDir
	oldChangesTmpDirPath := files.CreateTmpDir()
	MirrorTmpDirFilesFromTargetDir(config, targetDir, newChangesTmpDir, oldChangesTmpDirPath)
	fmt.Println(oldChangesTmpDirPath)

	// Copy tracked files missed to old dir, these are files not present in the new changes i.e. deleted in new changes
	CopyTrackedFiles(targetDir, oldChangesTmpDirPath)

	// Copy old config file if it exists
	oldConfigPath := targetDir + "/" + ConfigFileName
	_, err := os.Stat(oldConfigPath)
	if err == nil {
		CopyFile(oldConfigPath, oldChangesTmpDirPath+"/"+ConfigFileName)
	}

	// Write config changes to new changes
	WriteConfig(newChangesTmpDir, config)

	// Apply post actions
	ExecuteActions(newChangesTmpDir, config.Template.Actions.Post, templateContext, config)

	// Generate diff
	diff := git.Diff(oldChangesTmpDirPath, newChangesTmpDir)

	// Check there are changes
	if len(diff) == 0 {
		fmt.Println("No changes detected.")
		return true
	}

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
		git.Patch(targetDir, diff)
	case "d":
		// Dump patch to target dir
		var data = []byte(diff)
		var err = os.WriteFile(targetDir+"/patch.txt", data, 0664)
		interaction.HandleError(err, true)
	}

	// Update list of tracked files
	var trackedFiles []string
	err = filepath.Walk(newChangesTmpDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			var relativePath = path[len(newChangesTmpDir)+1:]
			trackedFiles = append(trackedFiles, relativePath)
		}
		return nil
	})
	interaction.HandleError(err, true)

	// Write tracked files to target dir's roller.yaml
	UpdateTrackedFiles(targetDir, trackedFiles, config)

	return true
}
