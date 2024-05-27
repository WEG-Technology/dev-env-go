package util

import (
	"bytes"
	"os/exec"
)

func RunCommand(cmd *exec.Cmd) (string, string, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	return out.String(), stderr.String(), err
}
