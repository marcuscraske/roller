package roller

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"roller/pkg/interaction"
)

// MirrorTmpDirFilesFromTargetDir copies files to outputDir from targetDir, where the same files exist in tmpDir
func MirrorTmpDirFilesFromTargetDir(config Config, targetDir string, tmpDir string, outputDir string) {

	fmt.Printf("Copying tracked files, outputDir=%s, tmpDir=%s, targetDir=%s\n", outputDir, tmpDir, targetDir)

	// Copy files from targetDir that exist in tmpDir
	err := filepath.Walk(tmpDir, func(tmpDirPath string, tmpDirInfo os.FileInfo, err error) error {
		relativePath := tmpDirPath[len(tmpDir):]

		targetDirPath := targetDir + relativePath
		outputDirPath := outputDir + relativePath

		if IsIgnoredFile(config, relativePath) {
			fmt.Println("Ignored file: " + relativePath)
		} else if !tmpDirInfo.IsDir() {
			// Copy the file if it exists in the target dir and it's a file
			targetDirInfo, err := os.Stat(targetDirPath)
			if !os.IsNotExist(err) && !targetDirInfo.IsDir() {
				CopyFile(targetDirPath, outputDirPath)
			}
		} else {
			// Create the folder...
			err = os.MkdirAll(outputDirPath, os.ModePerm)
			interaction.HandleError(err, true)
		}

		return nil
	})
	interaction.HandleError(err, true)
}

func CopyTrackedFiles(targetDir string, tmpDir string) {
	// Ensure tracked files in old project are copied across
	state, err := ReadState(targetDir)
	if err == nil {
		for _, relativePath := range state.TrackedFiles {
			targetDirPath := targetDir + "/" + relativePath
			tmpDirPath := tmpDir + "/" + relativePath

			// Check the file exists at the target dir, and then copy it
			_, err := os.Stat(targetDirPath)
			existsTargetDir := !os.IsNotExist(err)

			_, err = os.Stat(tmpDirPath)
			existsTmpDir := !os.IsNotExist(err)

			if existsTargetDir && !existsTmpDir {
				CopyFile(targetDirPath, tmpDirPath)
			}
		}
	}
}

func CopyFile(src string, dest string) {
	srcFile, err := os.Open(src)
	interaction.HandleError(err, true)
	defer srcFile.Close()

	newFile, err := os.Create(dest)
	interaction.HandleError(err, true)
	defer newFile.Close()

	_, err = io.Copy(newFile, srcFile)
	interaction.HandleError(err, true)

	fmt.Printf("Copied file, src=%s, dest=%s\n", src, dest)
}
