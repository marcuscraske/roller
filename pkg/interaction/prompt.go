package interaction

import (
	"fmt"
	"os"
	"strings"
)

func PromptOrExit(acceptedParams ...string) string {
	var answer string
	fmt.Scanln(&answer)

	answer = strings.ToLower(answer)

	for _, key := range acceptedParams {
		if answer == key {
			return answer
		}
	}

	fmt.Println("Aborted!")
	os.Exit(1)
	return ""
}
