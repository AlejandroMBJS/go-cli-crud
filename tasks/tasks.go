package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas")
		return
	}

	for _, task := range tasks {

		status := " "
		if task.Complete {
			status = "✓"
		} else {
			status = "○"
		}
		fmt.Printf("[%s] %d %s\n", status, task.ID, task.Name)
	}
}

func AddTask(tasks []Task, name string) []Task {
	newTask := Task{
		ID:       GetNextID(tasks),
		Name:     name,
		Complete: false,
	}
	return append(tasks, newTask)
}

func SaveTask(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		log.Printf("an error occurred: %s", err)
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Printf("an error occurred: %s", err)
	}
	err = file.Truncate(0)
	if err != nil {
		log.Printf("an error occurred: %s", err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		log.Printf("an error occurred: %s", err)
	}

	err = writer.Flush()
	if err != nil {
		log.Printf("an error occurred: %s", err)
	}
}

func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Complete = true
			break
		}
	}
	return tasks
}

func GetNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}
