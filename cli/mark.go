package main

import (
	"TaskTracker/json"
	"fmt"

	"github.com/urfave/cli/v2"
)

var mark = &cli.Command{
	Name:    "mark",
	Aliases: []string{"m"},
	Usage:   "mark task",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "id",
			Usage:    "id of task",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "status",
			Aliases:  []string{"s"},
			Usage:    "new status: todo, in-progress or done",
			Required: true,
		},
	},
	Action: markTask,
}

func markTask(ctx *cli.Context) error {
	id := ctx.Int("id")
	status := ctx.String("status")

	var marked bool
	if status == string(json.Todo) ||
		status == string(json.In_progress) ||
		status == string(json.Done) {
		marked = json.MarkTask(id, json.Status(status))

	} else {
		fmt.Println("you can choice only three statuses: todo, in-progress or done")
	}

	if marked {
		fmt.Println("successfully created!")
	} else {
		fmt.Println("task with current id not exists")
	}

	return nil
}
