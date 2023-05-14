package command

import (
	"os"
	"roller/pkg/roller"
)

func CommandCreate() bool {
	// TODO validate URL
	var gitUrl = os.Args[2]
	return roller.Create(gitUrl)
}
