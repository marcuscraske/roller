package roller

import (
	"container/list"
	"gopkg.in/yaml.v3"
	"os"
	"roller/pkg/interaction"
)

const ConfigFileName = "roller.yaml"

type Config struct {
	Template struct {
		Repo   string            `json:"repo"`
		Vars   map[string]string `json:"vars"`
		Ignore list.List         `json:"ignore"`
	} `json:"template"`
	Actions map[string]Action `json:"actions"`
}

// ReadConfig Read roller.yaml file.
func ReadConfig(dir string) Config {
	config := Config{}

	// Read file
	file, err := os.ReadFile(dir + "/" + ConfigFileName)
	interaction.HandleError(err)

	// Convert to yaml
	err = yaml.Unmarshal(file, &config)
	interaction.HandleError(err)

	// Apply enforced config
	config.Template.Ignore.PushFront(".git")

	return config
}

func MergeConfig(newConfig Config, oldConfig Config) Config {
	// TODO check if vars changed. if so, prompt to edit. write config either way...
	return newConfig
}
