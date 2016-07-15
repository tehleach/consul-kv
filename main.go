package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "Kyle Leach"
	app.Email = ""
	app.Usage = ""

	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
