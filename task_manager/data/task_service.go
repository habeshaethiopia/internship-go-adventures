package data

import (
	"encoding/json"
	// "fmt"
	"os"
	"task/models"
	"time"
)



var Tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}
// save the tasks to json file 

func SaveTasksToFile(tasks []models.Task, filename string) error {
    file, err := os.Create(filename)
	if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    return encoder.Encode(tasks)
}

// Load the tasks from the json file
func LoadTasksFromFile(filename string) ([]models.Task, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var tasks []models.Task
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&tasks)
    return tasks, err
}
