package main

import (
	"fmt"
	"os"

	"github.com/tehleach/consul-kv/command"
	"gopkg.in/urfave/cli.v1"
)

var path = cli.StringFlag{
	Name:  "prefix,p",
	Value: "",
	Usage: "Key prefix to access in consul kv",
}

var commands = []cli.Command{
	{
		Name:      "set",
		Usage:     "set a given kv",
		ArgsUsage: "<consulAddr> <key> <value>",
		Action:    command.CmdSet,
	},
	{
		Name:   "backup",
		Usage:  "backup to local file",
		Action: command.CmdBackup,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "from,f",
				Value: "",
				Usage: "Consul address to get values from, default localhost",
			},
			cli.StringFlag{
				Name:  "name,n",
				Value: "data.json",
				Usage: "Filename to save to",
			},
			path,
		},
	},
	{
		Name:   "restore",
		Usage:  "restore kvs from file/url to consul",
		Action: command.CmdRestore,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "from,f",
				Value: "",
				Usage: "Consul address to get values from, default localhost",
			},
			cli.StringFlag{
				Name:  "name,n",
				Value: "data.json",
				Usage: "Filename get values from",
			},
			cli.StringFlag{
				Name:  "to,t",
				Value: "",
				Usage: "Consul address to write values to, default localhost",
			},
			path,
		},
	},
}

func commandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
