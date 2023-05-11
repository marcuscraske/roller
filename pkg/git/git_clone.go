package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
	"roller/pkg/interaction"
)

// GitClone Clones the URL to a new tmp dir, and returns the path of the tmp dir.
func GitClone(url string) string {
	// TODO ability to handle private key
	// Create tmp folder for clone
	dir, err := os.MkdirTemp("", "roller")
	interaction.HandleError(err)

	// Clone the repo
	fmt.Printf("Cloning repo, tmpdir=%s, url=%s\n", dir, url)
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	interaction.HandleError(err)

	return dir
}
