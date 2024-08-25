package json

import (
	"slices"
	"time"
)

func AddTask(description string) {
	var ID int
	tasks := read()
	if len(tasks) == 0 {
		ID = 1
	} else {
		ID = tasks[len(tasks)-1].ID
	}
	task := newTask(ID, description)
	tasks = append(tasks, task)
	write(tasks)
}

func UpdateTask(ID int, description string) bool {
	tasks := read()
	index := slices.IndexFunc(tasks, func(task Task) bool {
		return task.ID == ID
	})
	if index == -1 {
		return false
	} else {
		tasks[index].Description = description
		tasks[index].UpdatedAt = time.Now()
		write(tasks)
		return true
	}
}

func DeleteTask(ID int) bool {
	tasks := read()
	var updated bool
	tasks = slices.DeleteFunc(tasks, func(task Task) bool {
		if task.ID == ID {
			updated = true
		}
		return task.ID == ID

	})
	write(tasks)
	return updated
}

func MarkTask(ID int, status Status) bool {
	tasks := read()
	index := slices.IndexFunc(tasks, func(task Task) bool {
		return task.ID == ID
	})
	if index == -1 {
		return false
	} else {
		tasks[index].Status = status
		tasks[index].UpdatedAt = time.Now()
		write(tasks)
		return true
	}
}

func ListTasks() Tasks {
	return read()
}
func ListTasksByStatus(status Status) Tasks {
	tasks := read()

	var newTasks Tasks

	for _, task := range tasks {
		if task.Status == status {
			newTasks = append(newTasks, task)
		}
	}
	return newTasks
}
