package command

import (
	"roller/pkg/interaction"
	"roller/pkg/roller"
)

func CommandCreate() bool {
	args := interaction.ParseAppArgs()
	var gitUrl = args["2"]
	var gitReference = args["3"]
	return roller.Create(gitUrl, gitReference)
}
