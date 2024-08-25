package main

import (
	"TaskTracker/json"
	"fmt"

	"github.com/urfave/cli/v2"
)

var list = &cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "shows task list",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "status",
			Aliases: []string{"s"},
			Value:   "none",
			Usage:   "shows list with status: todo, in-progress or done",
		},
	},
	Action: showList,
}

func showList(ctx *cli.Context) error {
	var tasks json.Tasks
	status := ctx.String("status")
	if status == "none" {
		tasks = json.ListTasks()
	} else if status == string(json.Todo) ||
		status == string(json.In_progress) ||
		status == string(json.Done) {
		tasks = json.ListTasksByStatus(json.Status(status))
	} else {
		fmt.Println("you can choice only three statuses: todo, in-progress or done")
	}

	for _, task := range tasks {
		fmt.Print(task.ID)
		fmt.Println(". " + task.Description + " (" + string(task.Status) + ")")

		fmt.Print("Created: ")
		fmt.Println(task.CreatedAt.Format("02 Jan 2006 15:04:05"))

		fmt.Print("Updated: ")
		fmt.Println(task.UpdatedAt.Format("02 Jan 2006 15:04:05"))

		fmt.Println()
	}
	return nil
}
