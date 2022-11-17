package storage

import "zut/internal/commands"

var Commands []*commands.Command = []*commands.Command{
	{
		Name:        "df",
		Description: "report file system disk space usage",
		Tags:        []string{"file", "disk", "space"},
	},
	{
		Name:        "rsync",
		Description: "a fast, versatile, remote (and local) file-copying tool",
		Tags:        []string{"copy"},
	},
}
