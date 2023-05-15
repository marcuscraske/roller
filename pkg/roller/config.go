package roller

import (
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"roller/pkg/interaction"
)

const ConfigFileName = "roller.yaml"

type Config struct {
	Template struct {
		Repo    string            `json:"repo"`
		Vars    map[string]string `json:"vars"`
		Replace map[string]string `json:"replace"`
		Ignore  []string          `json:"ignore"`
		Actions struct {
			Pre  []Action `json:"pre"`
			Post []Action `json:"post"`
		} `json:"actions"`
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
	result := newConfig

	isVarsSame := reflect.DeepEqual(newConfig.Template.Vars, oldConfig.Template.Vars)
	if !isVarsSame {
		result.Template.Vars = mergeStrMap(newConfig.Template.Vars, oldConfig.Template.Vars)
	}

	isReplaceSame := reflect.DeepEqual(newConfig.Template.Replace, oldConfig.Template.Replace)
	if !isReplaceSame {
		result.Template.Vars = mergeStrMap(newConfig.Template.Replace, oldConfig.Template.Replace)
	}

	isIgnoreSame := reflect.DeepEqual(newConfig.Template.Ignore, oldConfig.Template.Ignore)
	if !isIgnoreSame {
		result.Template.Ignore = mergeStrArray(newConfig.Template.Ignore, oldConfig.Template.Ignore)
	}

	return result
}

func mergeStrMap(new map[string]string, old map[string]string) map[string]string {
	result := map[string]string{}
	// Append all the new items
	for key, value := range new {
		result[key] = value
	}
	// Append/overwrite the old items
	for key, value := range old {
		result[key] = value
	}
	return result
}

func mergeStrArray(new []string, old []string) []string {
	var result []string
	keys := map[string]bool{}
	// Append all the new items
	for _, item := range new {
		result = append(result, item)
		keys[item] = true
	}
	// Append all the old items, providing they're not duplicates / already exist
	for _, item := range old {
		_, ok := keys[item]
		if !ok {
			keys[item] = true
			result = append(result, item)
		}
	}
	return result
}
