package json

import "time"

type Status string

const (
	Todo        Status = "todo"
	In_progress Status = "in-progress"
	Done        Status = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Tasks []Task

func newTask(ID int, description string) Task {
	return Task{
		ID:          ID + 1,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
