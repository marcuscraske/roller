package main

import "fmt"

func CommandHelp() {
	fmt.Println("Usage: roller COMMAND")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println(" create")
	fmt.Println(" update")
	fmt.Println(" version")
	fmt.Println("")
	fmt.Println("To get more help with roller, checkout TBC.")
}
