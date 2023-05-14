package command

import (
	"fmt"
)

var BuildVersion = "dev"
var BuildDate = "none"

func CommandVersion() bool {
	fmt.Printf("Version: %s\n", BuildVersion)
	fmt.Printf("Built:   %s\n", BuildDate)
	return true
}
