package roller

import (
	"container/list"
	"gopkg.in/yaml.v3"
	"os"
	"roller/pkg/interaction"
)

type Config struct {
	Template struct {
		Repo         string            `json:"repo"`
		Vars         map[string]string `json:"vars"`
		Ignore       list.List         `json:"ignore"`
		TrackedFiles []string          `json:"tracked_files"`
	} `json:"template"`
	Actions map[string]Action `json:"actions"`
}

// ReadConfig Read roller.yaml file.
func ReadConfig(path string) Config {
	config := Config{}

	// Read file
	file, err := os.ReadFile(path)
	interaction.HandleError(err)

	// Convert to yaml
	err = yaml.Unmarshal(file, &config)
	interaction.HandleError(err)

	// Apply enforced config
	config.Template.Ignore.PushFront(".git")
	config.Template.Ignore.PushFront("roller.yaml")

	return config
}
