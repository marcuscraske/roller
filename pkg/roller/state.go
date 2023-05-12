package roller

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"roller/pkg/interaction"
)

const StateFileName = ".roller.state"

type State struct {
	TrackedFiles []string `json:"tracked_files"`
}

func ReadState(dir string) State {
	state := State{}

	_, err := os.Stat(dir + "/.roller.state")
	if !errors.Is(err, os.ErrNotExist) {
		file, err := os.ReadFile(dir + "/" + StateFileName)
		interaction.HandleError(err)

		err = yaml.Unmarshal(file, &state)
		interaction.HandleError(err)
	}

	return state
}

func WriteState(dir string, state State) {
	data, err := yaml.Marshal(state)
	interaction.HandleError(err)

	err = os.WriteFile(dir+"/"+StateFileName, data, 0664)
	interaction.HandleError(err)
}
