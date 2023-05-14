package command

import "fmt"

func CommandVersion() bool {
	fmt.Println("1.0.0")
	return true
}
