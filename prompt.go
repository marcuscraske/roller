package main

import (
	"fmt"
	"os"
)

func PromptYesOrExit() {
	var answer string
	fmt.Scanln(&answer)
	if len(answer) == 0 || (answer != "y" && answer != "yes") {
		fmt.Println("Expected 'y' or 'yes', aborted!")
		os.Exit(1)
	}
}
