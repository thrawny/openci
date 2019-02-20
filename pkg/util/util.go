package util

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func RunCmd(cmd *exec.Cmd) (string, error) {
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		return strings.TrimSpace(errb.String()), err
	}
	return strings.TrimSpace(outb.String()), err
}
