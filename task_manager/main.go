package main

import (
	"fmt"
	"github.com/fatih/color"
	"task/data"
	"task/router"
)


func main() {
	fmt.Println("Task Manager API")
	tasks, err := data.LoadTasksFromFile("tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks from file", err)
	}
	data.Tasks = tasks
	r := router.Router()
	color.Green("Server is running on port 8080")
	color.Red("Read From file is Success")
	// color.Red("readfrom file")
	r.Run()
}