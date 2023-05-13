package roller

import (
	"strings"
)

func IsIgnoredFile(config Config, relativePath string) bool {

	ignored := false

	// Check whether the path matches any ignored files
	for _, val := range config.Template.Ignore {
		index := strings.Index(relativePath, val)
		if index >= 0 {
			ignored = true
		}
	}

	return ignored
}
