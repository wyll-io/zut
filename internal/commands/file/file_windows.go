//go:build windows
// +build windows

package file

import "zut/internal/commands"

var Commands []*commands.Command = []*commands.Command{
	{
		Name:        "copy",
		Description: "copy files and directories",
		Tags:        []string{},
		Examples:    []string{"copy source destination"},
	},
	{
		Name:        "dir",
		Description: "list files and directories",
		Tags:        []string{},
	},
}
