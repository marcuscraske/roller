package roller

import (
	"gopkg.in/yaml.v3"
	"os"
	"roller/pkg/interaction"
)

const ConfigFileName = "roller.yaml"

type Config struct {
	Template struct {
		Repo    string            `json:"repo"`
		Vars    map[string]string `json:"vars"`
		Replace map[string]string `json:"replace"`
		Ignore  []string          `json:"ignore"`
	} `json:"template"`
	Actions map[string]Action `json:"actions"`
}

// ReadConfig Read roller.yaml file.
func ReadConfig(dir string) (config Config, err error) {
	// Read file
	file, err := os.ReadFile(dir + "/" + ConfigFileName)
	interaction.HandleError(err, true)

	// Convert to yaml
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}

	// Apply enforced config
	config.Template.Ignore = append(config.Template.Ignore, ".git")
	return config, nil
}

func WriteConfig(dir string, config Config) {
	data, err := yaml.Marshal(config)
	interaction.HandleError(err, true)
	err = os.WriteFile(dir+"/"+ConfigFileName, data, 0664)
	interaction.HandleError(err, true)
}

// MergeConfig merges template config with target config, and provides the merged config
func MergeConfig(newConfig Config, oldConfig Config) Config {
	// TODO check config changed. if so, prompt to edit. write config either way...
	// TODO Template.Vars
	// TODO Template.Replace
	// TODO Template.Ignore
	return newConfig
}
