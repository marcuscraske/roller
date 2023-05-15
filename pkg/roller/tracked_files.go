package roller

import "roller/pkg/interaction"

func UpdateTrackedFiles(dir string, trackedFiles []string, config Config) {
	// Remove ignored files, from trackedFiles
	var result []string
	for _, path := range trackedFiles {
		if !IsIgnoredFile(config, path) {
			result = append(result, path)
		}
	}

	// Update state
	state, err := ReadState(dir)
	interaction.HandleError(err, true)

	state.TrackedFiles = result
	WriteState(dir, state)
}
