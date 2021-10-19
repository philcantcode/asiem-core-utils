package osutils

import (
	"asiem-core-utils/errutils"
	"os"
	"os/exec"
)

// Run BLAH
func Run(command string, args ...string) (string, bool) {
	resultByte, err := exec.Command(command, args...).CombinedOutput()

	if err != nil {
		return string(resultByte), false
	}

	return string(resultByte), true

}

func RunOnTop(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	errutils.ParseErr("Couldn't run on top: "+command, err)
}
