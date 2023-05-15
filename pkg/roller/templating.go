package roller

import (
	"fmt"
	"github.com/flosch/pongo2/v6"
	"os"
	"path/filepath"
	"roller/pkg/interaction"
	"strings"
)

func CreateTemplateContext(config Config) (templateContext pongo2.Context) {
	// Setup context
	context := pongo2.Context{}
	for key, value := range config.Template.Vars {
		context[key] = value
	}
	return context
}

// ApplyTemplatingDir Applies templating to the provided dir
func ApplyTemplatingDir(dir string, config Config, templateContext pongo2.Context) {
	// Iterate and apply template to each file
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return ApplyTemplatingFile(path, info, config, templateContext)
		}
		return nil
	})
	interaction.HandleError(err, true)
}

func ApplyTemplatingFile(path string, info os.FileInfo, config Config, templateContext pongo2.Context) error {
	// Read file content
	data, err := os.ReadFile(path)
	interaction.HandleError(err, true)

	output := string(data)

	// Apply templating
	output = ApplyTemplatingString(output, templateContext, config)

	// Write output to file
	data = []byte(output)
	err = os.WriteFile(path, data, info.Mode())
	interaction.HandleError(err, true)

	fmt.Printf("Template applied, file=%s\n", path)
	return nil
}

func ApplyTemplatingString(text string, templateContext pongo2.Context, config Config) string {
	// Apply template
	template, err := pongo2.FromString(text)
	if err != nil {
		fmt.Printf("Dump of failed template:\n%s\n", text)
	}
	interaction.HandleError(err, true)

	output, err := template.Execute(templateContext)
	interaction.HandleError(err, true)

	// Apply string replacement
	for key, value := range config.Template.Replace {
		output = strings.ReplaceAll(output, key, value)
	}

	return output
}
