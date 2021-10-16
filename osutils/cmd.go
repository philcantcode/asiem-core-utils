package osutils

import (
	"os/exec"
)

func Run(command string, args ...string) (result string, success bool) {

	resultByte, err := exec.Command(command, args...).CombinedOutput()

	if err != nil {
		return string(resultByte), false
	}

	return string(resultByte), true

}
