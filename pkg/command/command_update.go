package command

import "roller/pkg/roller"

func CommandUpdate() bool {
	args := ParseArgs()
	gitReference := args["2"]
	return roller.Update(gitReference)
}
