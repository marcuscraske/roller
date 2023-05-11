package main

import (
	"container/list"
	"gopkg.in/yaml.v3"
	"os"
)

type RollerConfig struct {
	template struct {
		repo          string
		vars          map[string]string
		ignore        list.List
		tracked_files []string
	}
	actions map[string]RollerAction
}

// ReadConfig Read roller.yaml file.
func ReadConfig(path string) RollerConfig {
	config := RollerConfig{}

	// Read file
	file, err := os.ReadFile(path)
	HandleError(err)

	// Convert to yaml
	err = yaml.Unmarshal(file, &config)
	HandleError(err)

	// Apply enforced config
	config.template.ignore.PushFront(".git")
	config.template.ignore.PushFront("roller.yaml")

	return config
}
