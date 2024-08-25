package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var app *cli.App = &cli.App{
	Name:  "TaskTracker",
	Usage: "cli task tracker",
	Commands: []*cli.Command{
		list,
		add,
		delete,
		update,
		mark,
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
