package main

import (
	"TaskTracker/json"
	"fmt"

	"github.com/urfave/cli/v2"
)

var add = &cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "add task",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "description",
			Aliases:  []string{"d"},
			Usage:    "description of new task",
			Required: true,
		},
	},
	Action: addTask,
}

func addTask(ctx *cli.Context) error {
	description := ctx.String("description")
	json.AddTask(description)
	fmt.Println("successfully created!")
	return nil
}
