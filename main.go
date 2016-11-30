package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {

	app := cli.NewApp()
	app.Author = "Kyle Leach"
	app.Email = ""
	app.Usage = ""

	app.Commands = commands
	app.CommandNotFound = commandNotFound

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
