package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	task "github.com/AlejandroMBJS/go-cli-crud/tasks"
)

func main() {

	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("an error occurred: %s", err)
	}
	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		log.Printf("an error occurred: %s", err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			log.Printf("an error occurred: %s", err)
		}
		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			log.Printf("an error occurred: %s", err)
		}

	} else {
		tasks = []task.Task{}
	}
	if len(os.Args) < 2 {
		printUsage()
	} else {

		switch os.Args[1] {
		case "list":
			task.ListTasks(tasks)
		case "add":
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Ingresa la nueva tarea")
			name, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("an error occurred: %s", err)
			}
			name = strings.TrimSpace(name)
			tasks = task.AddTask(tasks, name)
			task.SaveTask(file, tasks)
		case "delete":
			if len(os.Args) < 3 {
				fmt.Println("Debes proporciona ID | use list command to see all")
				return
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("El id Debe ser un solo numero")
				return
			}
			tasks = task.DeleteTask(tasks, id)
			task.SaveTask(file, tasks)
			fmt.Printf("Tarea %v Eliminada", id)
		case "complete":
			if len(os.Args) < 3 {
				fmt.Println("Debes proporciona ID | usa comando list para ver IDs")
				return
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("El id Debe ser un solo numero")
				return
			}
			tasks = task.CompleteTask(tasks, id)
			task.SaveTask(file, tasks)
			fmt.Printf("Tarea %v Completa", id)
		default:
			printUsage()
		}
	}
}

func printUsage() {
	fmt.Println("Uso: go-cli-crud [list|add|complete|delete]")
}
