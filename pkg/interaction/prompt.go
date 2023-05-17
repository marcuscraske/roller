package interaction

import (
	"bufio"
	"os"
	"strings"
)

func PromptOrBlank(acceptedParams ...string) string {
	answer := PromptReadLine()
	answer = strings.ToLower(answer)

	for _, key := range acceptedParams {
		if answer == key {
			return answer
		}
	}
	return ""
}

func PromptReadLine() string {
	result := ""
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		result = scanner.Text()
	}
	return result
}
