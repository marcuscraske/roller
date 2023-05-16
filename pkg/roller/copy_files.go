package roller

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"roller/pkg/interaction"
)

// MirrorTmpDirFilesFromTargetDir copies files to outputDir from targetDir, where the same files exist in tmpDir
func MirrorTmpDirFilesFromTargetDir(config Config, targetDir string, tmpDir string, outputDir string) {

	fmt.Printf("Copying tracked files, outputDir=%s, tmpDir=%s, targetDir=%s\n", outputDir, tmpDir, targetDir)

	// Copy files from targetDir that exist in tmpDir, to outputDir
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
			err = os.MkdirAll(outputDirPath, tmpDirInfo.Mode())
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
				err = createParentDirIfNotExist(targetDirPath, targetDirPath)
				interaction.HandleError(err, true)

				CopyFile(targetDirPath, tmpDirPath)
			}
		}
	}
}

// CreateParentDirIfNotExist Used to ensure the parent/containing directory for the given path exists, otherwise it's created
func createParentDirIfNotExist(filePath string, mirrorPath string) (err error) {
	// Check whether the parent folder exists
	dir := path.Dir(filePath)
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {

		// Fetch the file permissions of the mirror file
		mirrorDir := path.Dir(mirrorPath)
		mirrorDirInfo, err := os.Stat(mirrorDir)
		if err != nil {
			return err
		}

		err = os.MkdirAll(dir, mirrorDirInfo.Mode())
		return err
	}
	return nil
}

func CopyFile(src string, dest string) {
	// Read existing permissions
	info, err := os.Stat(src)
	interaction.HandleError(err, true)

	// Open source file
	srcFile, err := os.Open(src)
	interaction.HandleError(err, true)
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
			interaction.HandleError(err, true)
		}
	}(srcFile)

	// Create destination file and copy contents from source
	newFile, err := os.Create(dest)
	interaction.HandleError(err, true)
	defer func(newFile *os.File) {
		err := newFile.Close()
		if err != nil {
			interaction.HandleError(err, true)
		}
	}(newFile)

	_, err = io.Copy(newFile, srcFile)
	interaction.HandleError(err, true)

	// Set new file's permissions to match the source
	err = os.Chmod(dest, info.Mode())
	interaction.HandleError(err, true)

	fmt.Printf("Copied file, src=%s, dest=%s, mode=%d\n", src, dest, info.Mode())
}
