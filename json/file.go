package json

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func init() {
	if _, err := os.Stat("tasks.json"); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create("tasks.json")
		if err != nil {
			log.Fatal(err)
		}
		file.Write([]byte("{}"))
		file.Close()
	}
}

func read() Tasks {
	var tasks Tasks
	file, err := os.ReadFile("tasks.json")

	if err != nil {
		log.Fatal(err)
	}

	if string(file) == "{}" {
		return tasks
	}

	err = json.Unmarshal(file, &tasks)

	if err != nil {
		log.Fatal(err)
	}

	return tasks
}

func write(tasks Tasks) {
	file, err := os.Create("tasks.json")

	if err != nil {
		log.Fatal(err)
	}

	jsonTasks, err := json.MarshalIndent(tasks, " ", " ")

	if err != nil {
		log.Fatal(err)
	}

	file.Write(jsonTasks)
}
