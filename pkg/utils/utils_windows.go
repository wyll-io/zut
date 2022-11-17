//go:build windows
// +build windows

package utils

func CommandExists(cmd string) bool {
	return true
}
