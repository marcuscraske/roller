package command

import (
	"roller/pkg/interaction"
	"roller/pkg/roller"
)

func CommandUpdate() bool {
	args := interaction.ParseAppArgs()
	gitReference := args["2"]
	return roller.Update(gitReference)
}
