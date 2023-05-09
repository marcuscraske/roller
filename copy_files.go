package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateTmpDirAndCopyTrackedFiles(config RollerConfig, targetDir string, tmpDir string) string {
	// Create "old" tmp dir to copy files that already exist in the target dir
	oldChangesTmpDirPath, err := os.MkdirTemp("", "roller")
	HandleError(err)

	fmt.Printf("Copying tracked files, oldChangesTmpDirPath=%s, tmpDir=%s, targetDir=%s\n", oldChangesTmpDirPath, tmpDir, targetDir)

	// Copy files from targetDir that exist in tmpDir
	err = filepath.Walk(tmpDir, func(tmpDirPath string, tmpDirInfo os.FileInfo, err error) error {
		relativePath := tmpDirPath[len(tmpDir):]

		targetDirPath := targetDir + relativePath
		oldChangesTmpDirPath := oldChangesTmpDirPath + relativePath

		if IsIgnoredFile(config, relativePath) {
			fmt.Println("Ignored file: " + relativePath)
		} else if !tmpDirInfo.IsDir() {
			// Copy the file if it exists in the target dir and it's a file
			targetDirInfo, err := os.Stat(targetDirPath)
			if !os.IsNotExist(err) && !targetDirInfo.IsDir() {
				CopyFile(targetDirPath, oldChangesTmpDirPath)
			}
		} else {
			// Create the folder...
			err = os.MkdirAll(oldChangesTmpDirPath, os.ModePerm)
			HandleError(err)
		}

		return nil
	})
	HandleError(err)

	return oldChangesTmpDirPath
}

func CopyFile(src string, dest string) {
	srcFile, err := os.Open(src)
	HandleError(err)
	defer srcFile.Close()

	newFile, err := os.Create(dest)
	HandleError(err)
	defer newFile.Close()

	_, err = io.Copy(newFile, srcFile)
	HandleError(err)

	fmt.Printf("Copied file, src=%s, dest=%s\n", src, dest)
}
