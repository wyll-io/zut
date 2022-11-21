//go:build windows
// +build windows

package utils

func CommandIsInstalled(cmd string) bool {
	return true
}
