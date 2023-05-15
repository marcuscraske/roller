package roller

import (
	"fmt"
	"github.com/flosch/pongo2/v6"
	"os"
	"path/filepath"
	"roller/pkg/interaction"
	"strings"
	"unicode/utf8"
)

func CreateTemplateContext(config Config) (templateContext pongo2.Context) {
	// Setup context
	context := pongo2.Context{}
	for key, value := range config.Template.Vars {
		context[key] = value.Value
	}
	return context
}

// ApplyTemplatingDir Applies templating to the provided dir
func ApplyTemplatingDir(dir string, config Config, templateContext pongo2.Context, applyReplace bool) {
	// Iterate and apply template to each file
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			relativePath := path[len(dir)+1:]
			return ApplyTemplatingFile(path, relativePath, info, config, templateContext, applyReplace)
		}
		return nil
	})
	interaction.HandleError(err, true)
}

func ApplyTemplatingFile(path string, relativePath string, info os.FileInfo, config Config, templateContext pongo2.Context, applyReplace bool) error {
	// Ignore roller.yaml
	// TODO support excluded files from templating
	if relativePath == ConfigFileName {
		return nil
	}

	// Read file content
	data, err := os.ReadFile(path)
	interaction.HandleError(err, true)

	output := string(data)

	// Check file is valid text content, otherwise abort
	if !utf8.ValidString(output) {
		fmt.Printf("Ignored file for templating as invalid utf8, path=%s\n", path)
		return nil
	}

	// Apply templating
	output, err = ApplyTemplatingString(output, templateContext, config, applyReplace)
	if err != nil {
		fmt.Printf("WARNING: failed to apply templating, path=%s, error=%s\n", path, err)
	} else {
		// Write output to file
		data = []byte(output)
		err = os.WriteFile(path, data, info.Mode())
		interaction.HandleError(err, true)
	}

	fmt.Printf("Template applied, file=%s\n", path)
	return nil
}

func ApplyTemplatingString(text string, templateContext pongo2.Context, config Config, applyReplace bool) (output string, err error) {
	output = text

	// Apply string replacement
	if applyReplace {
		for key, value := range config.Template.Replace {
			// Apply templating to key/values
			key, _ = ApplyTemplatingString(key, templateContext, config, false)
			value, _ = ApplyTemplatingString(value, templateContext, config, false)

			// Perform string replacement
			output = strings.ReplaceAll(output, key, value)
		}
	}

	// Apply template
	template, err := pongo2.FromString(output)
	if err != nil {
		return output, err
	}

	output, err = template.Execute(templateContext)
	if err != nil {
		return output, err
	}

	return output, nil
}
