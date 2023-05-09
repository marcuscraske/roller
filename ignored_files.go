package main

import (
	"strings"
)

func IsIgnoredFile(config RollerConfig, relativePath string) bool {

	ignored := false

	// Check whether the path matches any ignored files
	for e := config.template.ignore.Front(); e != nil; e = e.Next() {
		val := e.Value.(string)
		index := strings.Index(relativePath, val)
		if index >= 0 {
			ignored = true
		}
	}

	return ignored
}
