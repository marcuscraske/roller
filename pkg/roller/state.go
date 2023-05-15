package roller

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"roller/pkg/interaction"
)

const StateFileName = ".roller.state.yaml"

type State struct {
	TrackedFiles []string `json:"tracked_files"`
}

func ReadState(dir string) (state State, err error) {
	_, err = os.Stat(dir + "/.roller.state.yaml")
	if !errors.Is(err, os.ErrNotExist) {
		file, err := os.ReadFile(dir + "/" + StateFileName)
		if err != nil {
			return state, err
		}

		err = yaml.Unmarshal(file, &state)
		if err != nil {
			return state, err
		}
	}

	return state, nil
}

func WriteState(dir string, state State) {
	data, err := yaml.Marshal(state)
	interaction.HandleError(err, true)

	err = os.WriteFile(dir+"/"+StateFileName, data, 0664)
	interaction.HandleError(err, true)
}
