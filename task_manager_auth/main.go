package main

import (
	"fmt"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"

	// "task/data"
	"task/data"
	"task/router"
)


func main() {
	fmt.Println("Task Manager API")
	var client *mongo.Client 
	data.Tasks, client=data.CreateDB()
	
	r := router.Router()
	color.Green("Server is running on port 8080")
	r.Run()
	data.CloseDB(client)
}