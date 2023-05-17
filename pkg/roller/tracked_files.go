package roller

func UpdateTrackedFiles(state State, trackedFiles []string, config Config) {
	// Remove ignored files, from trackedFiles
	var result []string
	for _, path := range trackedFiles {
		if !IsIgnoredFile(config, path) {
			result = append(result, path)
		}
	}
	state.TrackedFiles = result
}
