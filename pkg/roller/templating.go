package roller

import (
	"fmt"
	"github.com/flosch/pongo2/v6"
	"os"
	"path/filepath"
	"roller/pkg/interaction"
	"strings"
)

// ApplyTemplating Applies templating to the provided dir
func ApplyTemplating(dir string, config Config) {
	// Read config for vars
	config, err := ReadConfig(dir)

	// Setup context
	context := pongo2.Context{}
	for key, value := range config.Template.Vars {
		context[key] = value
	}

	// Iterate and apply template to each file
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return ApplyTemplatingFile(path, info, config, context)
		}
		return nil
	})
	interaction.HandleError(err)
}

func ApplyTemplatingFile(path string, info os.FileInfo, config Config, context pongo2.Context) error {
	// Apply templating
	template, err := pongo2.FromFile(path)
	interaction.HandleError(err)

	output, err := template.Execute(context)
	interaction.HandleError(err)

	// Apply string replacement
	for key, value := range config.Template.Replace {
		output = strings.ReplaceAll(output, key, value)
	}

	// Write output to file
	data := []byte(output)
	err = os.WriteFile(path, data, info.Mode())
	interaction.HandleError(err)

	fmt.Printf("Template applied, file=%s\n", path)
	return nil
}
