package roller

func UpdateTrackedFiles(dir string, trackedFiles []string, config Config) {
	// Remove ignored files, from trackedFiles
	// TODO...

	// Update state
	var state = ReadState(dir)
	state.TrackedFiles = trackedFiles
	WriteState(dir, state)
}
