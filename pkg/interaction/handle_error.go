package interaction

import (
	"fmt"
	"os"
)

// HandleError Checks whether the error is not nil; if so, it's printed to stdout and the program exits
func HandleError(err error) {
	// TODO use panic(err) instead?
	if err != nil {
		// Dump the error
		fmt.Printf("Failed due to unexpected error: %s", err)

		// Hard exit with error-code...
		os.Exit(1)
	}
}
