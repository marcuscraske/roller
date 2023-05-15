package roller

import "roller/pkg/interaction"

func UpdateTrackedFiles(dir string, trackedFiles []string, config Config) {
	// Remove ignored files, from trackedFiles
	// TODO is this actually needed?

	// Update state
	state, err := ReadState(dir)
	interaction.HandleError(err, true)

	state.TrackedFiles = trackedFiles
	WriteState(dir, state)
}
