package git

import (
	"fmt"
	"roller/pkg/files"
	"roller/pkg/interaction"
)

// Clone Clones the URL to a new tmp dir, and returns the path of the tmp dir.
func Clone(url string) string {
	// Create tmp folder for clone
	dir := files.CreateTmpDir()

	// Clone the repo, launch it as an interactive process for auth
	fmt.Printf("Cloning repo, tmpdir=%s, url=%s\n", dir, url)
	err := interaction.LaunchInteractiveProcess(dir, "git", "clone", url, ".")
	if err != nil {
		fmt.Println("Failed to clone repository, aborted!")
		interaction.HandleError(err, false)
	}

	return dir
}
