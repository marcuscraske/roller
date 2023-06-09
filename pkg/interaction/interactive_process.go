package interaction

import (
	"fmt"
	"os"
	"os/exec"
)

func LaunchInteractiveProcess(workingDir string, name string, arg ...string) (err error) {
	fmt.Printf("Launching interactive process, name=%s, args=%s\n", name, arg)
	process := exec.Command(name, arg...)
	process.Dir = workingDir
	stdin, err := process.StdinPipe()
	HandleError(err, true)

	defer stdin.Close()
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	if err = process.Start(); err != nil {
		return err
	}

	err = process.Wait()
	return err
}
