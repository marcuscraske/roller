package roller

import (
	"fmt"
	"github.com/flosch/pongo2/v6"
	"roller/pkg/interaction"
	"strings"
)

type Action struct {
	Shell       string `yaml:"shell"`
	WorkingDir  string `yaml:"dir"`
	IgnoreError bool   `yaml:"ignore"`
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
	shell, err := ApplyTemplatingString(shell, templateContext, config, false)
	interaction.HandleError(err, true)

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
	err = interaction.LaunchInteractiveProcess(workingDir, name, args...)
	if err != nil {
		if !action.IgnoreError {
			interaction.HandleError(err, false)
		} else {
			fmt.Printf("WARNING: failed to execute action - name=%s, args=%s, error=%s\n", name, args, err)
		}
	}
	return true
}
