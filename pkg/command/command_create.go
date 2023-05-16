package command

import (
	"roller/pkg/roller"
)

func CommandCreate() bool {
	args := ParseArgs()
	var gitUrl = args["2"]
	var gitReference = args["3"]
	return roller.Create(gitUrl, gitReference)
}
