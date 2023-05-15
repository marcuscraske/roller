package roller

import (
	"github.com/flosch/pongo2/v6"
	"roller/pkg/interaction"
	"strings"
)

type Action struct {
	Shell      string `json:"shell"`
	WorkingDir string `json:"working_dir"`
}

func ExecuteActionByName(defaultWorkingDir string, actionName string, templateContext pongo2.Context, config Config) bool {
	// Pull the action and execute it, if it exists
	var action, exists = config.Actions[actionName]
	if !exists {
		return false
	}
	return ExecuteAction(defaultWorkingDir, action, templateContext, config)
}

func ExecuteActions(defaultWorkingDir string, actions []Action, templateContext pongo2.Context, config Config) {
	for _, action := range actions {
		ExecuteAction(defaultWorkingDir, action, templateContext, config)
	}
}

func ExecuteAction(defaultWorkingDir string, action Action, templateContext pongo2.Context, config Config) bool {

	// Apply templating
	shell := action.Shell
	shell = ApplyTemplatingString(shell, templateContext, config)

	// Determine name and args
	var parts = strings.Split(shell, " ")
	var name = parts[0]
	var args []string
	if len(parts) > 1 {
		args = parts[1:]
	}

	// Determine working directory
	workingDir := action.WorkingDir
	if len(workingDir) == 0 {
		workingDir = defaultWorkingDir
	}

	// Launch the process
	err := interaction.LaunchInteractiveProcess(workingDir, name, args...)
	interaction.HandleError(err, false)
	return true
}
