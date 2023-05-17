package interaction

import (
	"fmt"
	"strings"
)

func PromptOrBlank(acceptedParams ...string) string {
	var answer string
	fmt.Scanln(&answer)

	answer = strings.ToLower(answer)

	for _, key := range acceptedParams {
		if answer == key {
			return answer
		}
	}
	return ""
}
