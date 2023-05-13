package roller

import (
	"fmt"
	"os"
	"path/filepath"
	"roller/pkg/interaction"
)

func TemplatingApply(dir string, config Config) {
	// TBC...
}

// ApplyTemplating Applies templating to the provided dir
func ApplyTemplating(dir string) {
	// Read config for vars
	config, err := ReadConfig(dir)

	// Setup context
	//context := pongo2.Context{"vars": config.template.vars}

	// Iterate and apply template to each file
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		return ApplyTemplatingFile(dir, path, info, config)
	})
	interaction.HandleError(err)
}

func ApplyTemplatingFile(dir string, path string, info os.FileInfo, config Config) error {
	//if IsIgnoredFile(config, dir, path, info) {
	//	fmt.Printf("Skipped templating, file=%s\n", path)
	//	return nil
	//}

	fmt.Printf("Processing, file=%s\n", path)

	// Apply templating

	// Check whether new/diff to target file

	return nil
}
