package command

import (
	"fmt"
)

var BuildVersion = "dev"
var BuildCommit = "n/a"
var BuildDate = "none"

func CommandVersion() bool {
	fmt.Printf("Version: %s\n", BuildVersion)
	fmt.Printf("Commit:  %s\n", BuildCommit)
	fmt.Printf("Built:   %s\n", BuildDate)
	return true
}
