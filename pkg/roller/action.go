package roller

import (
	"os"
	"roller/pkg/interaction"
	"strings"
)

type Action struct {
	Shell      string `json:"shell"`
	WorkingDir string `json:"working_dir"`
}

func ExecuteAction(actionName string) bool {
	var targetDir, err = os.Getwd()
	interaction.HandleError(err, true)

	// Read config and check the action exists
	var config, err2 = ReadConfig(targetDir)
	interaction.HandleError(err2, true)

	var action, exists = config.Actions[actionName]
	if !exists {
		return false
	}

	var parts = strings.Split(action.Shell, " ")
	var name = parts[0]
	var args []string
	if len(parts) > 1 {
		args = parts[1:]
	}

	var workingDir = action.WorkingDir
	if len(workingDir) == 0 {
		workingDir, err = os.Getwd()
		interaction.HandleError(err, true)
	}

	err = interaction.LaunchInteractiveProcess(workingDir, name, args...)
	interaction.HandleError(err, false)
	return true
}
