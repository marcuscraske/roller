package command

import (
	"fmt"
	"os"
	"roller/pkg/interaction"
	"roller/pkg/roller"
)

func CommandHelp() bool {
	fmt.Println("Usage: roller COMMAND")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println(" create")
	fmt.Println(" update")
	fmt.Println(" version")
	fmt.Println("")

	targetDir, err := os.Getwd()
	interaction.HandleError(err, true)

	config, err := roller.ReadConfig(targetDir)
	if err == nil {
		if len(config.Actions) > 0 {
			fmt.Println("Available custom commands:")
			for key := range config.Actions {
				fmt.Printf("  %s\n", key)
			}
			fmt.Println("")
		}
	}

	fmt.Println("To get more help with roller, checkout TBC.")
	return true
}
