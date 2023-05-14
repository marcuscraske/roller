package interaction

import (
	"fmt"
	"runtime/debug"
)

// HandleError Checks whether the error is not nil; if so, it's printed to stdout and the program exits
func HandleError(err error, critical bool) {
	if err != nil {
		fmt.Printf("Failed due to unexpected error: %s\n", err)
		if critical {
			debug.PrintStack()
		}
		panic(err)
	}
}
