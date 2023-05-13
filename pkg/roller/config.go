package roller

import (
	"gopkg.in/yaml.v3"
	"os"
	"roller/pkg/interaction"
)

const ConfigFileName = "roller.yaml"

type Config struct {
	Template struct {
		Repo   string            `json:"repo"`
		Vars   map[string]string `json:"vars"`
		Ignore []string          `json:"ignore"`
	} `json:"template"`
	Actions map[string]Action `json:"actions"`
}

// ReadConfig Read roller.yaml file.
func ReadConfig(dir string) (config Config, err error) {
	// Read file
	file, err := os.ReadFile(dir + "/" + ConfigFileName)
	interaction.HandleError(err)

	// Convert to yaml
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}

	// Apply enforced config
	config.Template.Ignore = append(config.Template.Ignore, ".git")
	return config, nil
}

func MergeConfig(newConfig Config, oldConfig Config) Config {
	// TODO check if vars changed. if so, prompt to edit. write config either way...
	return newConfig
}
