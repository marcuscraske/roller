package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
)

// GitClone Clones the URL to a new tmp dir, and returns the path of the tmp dir.
func GitClone(url string) string {
	// TODO ability to handle private key
	// Create tmp folder for clone
	dir, err := os.MkdirTemp("", "roller")
	HandleError(err)

	// Clone the repo
	fmt.Printf("Cloning repo, tmpdir=%s, url=%s\n", dir, url)
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	HandleError(err)

	return dir
}
