package interaction

import (
	"fmt"
	"os"
	"os/exec"
)

func LaunchInteractiveProcess(name string, arg ...string) {
	fmt.Println("Launching vim for roller.yaml...")
	process := exec.Command(name, arg...)
	stdin, err := process.StdinPipe()
	HandleError(err)

	defer stdin.Close()
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	if err = process.Start(); err != nil {
		HandleError(err)
	}

	err = process.Wait()
	if err != nil {
		HandleError(err)
	}
}
