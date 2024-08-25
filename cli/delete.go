package main

import (
	"TaskTracker/json"
	"fmt"

	"github.com/urfave/cli/v2"
)

var delete = &cli.Command{
	Name:    "delete",
	Aliases: []string{"d"},
	Usage:   "delete task",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "id",
			Usage:    "id of the task to delete",
			Required: true,
		},
	},
	Action: deleteTask,
}

func deleteTask(ctx *cli.Context) error {
	id := ctx.Int("id")

	deleted := json.DeleteTask(id)

	if deleted {
		fmt.Println("successfully deleted!")
	} else {
		fmt.Println("task with current id not exists")
	}
	return nil
}
