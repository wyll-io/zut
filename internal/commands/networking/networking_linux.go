//go:build linux
// +build linux

package networking

import "zut/internal/commands"

var Commands []*commands.Command = []*commands.Command{
	{
		Name:        "ifconfig",
		Description: "",
		Tags:        []string{"ip"},
	},
	{
		Name:        "traceroute",
		Description: "",
		Tags:        []string{"ip"},
	},
	{
		Name:        "tcpdump",
		Description: "",
		Tags:        []string{"ip"},
	},
	{
		Name:        "netstat",
		Description: "",
		Tags:        []string{"ip"},
		Examples:    []string{"netstat -anlp"},
	},
}
