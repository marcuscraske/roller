package files

import (
	"fmt"
	"os"
	"roller/pkg/interaction"
)

var dirs []string

func CreateTmpDir() string {
	tmpDir, err := os.MkdirTemp("", "roller")
	interaction.HandleError(err)
	dirs = append(dirs, tmpDir)
	fmt.Printf("Created tmpdir=%s\n", tmpDir)
	return tmpDir
}

func TmpDirCleanup() {
	for _, path := range dirs {
		err := os.RemoveAll(path)
		if err == nil {
			fmt.Printf("Cleaned tmpdir=%s\n", path)
		} else {
			fmt.Printf("WARNING: failed to remove tmpdir, %s\n", err.Error())
		}
	}
}
