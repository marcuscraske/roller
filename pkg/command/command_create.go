package command

import (
	"os"
	"roller/pkg/roller"
)

func CommandCreate() {
	// TODO validate URL
	var gitUrl = os.Args[2]
	roller.Create(gitUrl)
}
