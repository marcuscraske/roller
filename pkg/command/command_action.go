package command

import (
	"os"
	"roller/pkg/roller"
)

func CommandAction() bool {
	if len(os.Args) < 2 {
		return false
	}
	var action = os.Args[1]
	return roller.ExecuteAction(action)
}
