package command

import "roller/pkg/roller"

func CommandUpdate() bool {
	return roller.Update()
}
