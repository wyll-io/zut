//go:build linux
// +build linux

package utils

import (
	"errors"
	"os/exec"
)

func CommandIsInstalled(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if errors.Is(err, exec.ErrDot) {
		return true
	}

	return err == nil
}
