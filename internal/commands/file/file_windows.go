//go:build windows
// +build windows

package file

import "zut/internal/commands"

var Commands []*commands.Command = []*commands.Command{
	{
		Name:        "copy",
		Description: "copy files and directories",
		Tags:        []string{"copy", "files", "directories"},
	},
	{
		Name:        "dir",
		Description: "list files and directories",
		Tags:        []string{"list", "files", "directories"},
	},
}
