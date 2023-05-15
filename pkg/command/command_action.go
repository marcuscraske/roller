package command

import (
	"os"
	"roller/pkg/interaction"
	"roller/pkg/roller"
)

func CommandAction() bool {
	if len(os.Args) < 2 {
		return false
	}

	// Read config
	targetDir, err := os.Getwd()
	interaction.HandleError(err, true)

	config, err := roller.ReadConfig(targetDir)
	interaction.HandleError(err, true)

	// Setup template context
	templateContext := roller.CreateTemplateContext(config)

	// Execute action
	var action = os.Args[1]
	return roller.ExecuteActionByName(targetDir, action, templateContext, config)
}
