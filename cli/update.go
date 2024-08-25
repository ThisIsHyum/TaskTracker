package main

import (
	"TaskTracker/json"
	"fmt"

	"github.com/urfave/cli/v2"
)

var update = &cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Usage:   "update task",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "id",
			Usage:    "id of task",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "description",
			Aliases:  []string{"d"},
			Usage:    "new description of task",
			Required: true,
		},
	},
	Action: updateTask,
}

func updateTask(ctx *cli.Context) error {
	id := ctx.Int("id")
	description := ctx.String("description")
	updated := json.UpdateTask(id, description)

	if updated {
		fmt.Println("successfully updated!")
	} else {
		fmt.Println("task with current id not exists")
	}
	return nil
}
