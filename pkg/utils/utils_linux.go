//go:build linux
// +build linux

package utils

import (
	"errors"
	"os/exec"
)

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if errors.Is(err, exec.ErrDot) {
		err = nil
	}

	return err == nil
}
