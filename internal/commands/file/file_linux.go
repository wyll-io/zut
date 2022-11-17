//go:build linux
// +build linux

package file

import "zut/internal/commands"

var Commands []*commands.Command = []*commands.Command{
	{
		Name:        "cp",
		Description: "copy files and directories",
		Tags:        []string{"copy", "files", "directories"},
	},
	{
		Name:        "rsync",
		Description: "a fast, versatile, remote (and local) file-copying tool",
		Tags:        []string{"copy"},
		Examples:    []string{"rsync -ah --progress source destination"},
	},
	{
		Name:        "pwd",
		Description: "print name of current/working directory",
	},
}
