package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	// "go.mongodb.org/mongo-driver/mongo"
	"github.com/joho/godotenv"

	"task/data"
	"task/router"
)

func main() {
	fmt.Println("Task Manager API")
	
	err := godotenv.Load()
	if err != nil {
		color.Red("Error loading .env file : %v",err  )
	}
	 //!2 Access the environment variable
	 jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	 if jwtSecretKey == "" {
		 fmt.Println("JWT_SECRET_KEY is not set")
	 } else {
		 fmt.Println("JWT_SECRET_KEY:", jwtSecretKey)
	 }
	
	client, err := data.ConnectDB()
	if err != nil {
		color.Red("Error connecting to database")
	}
	data.Users = data.CreateCollection(client, "task_manager", "users")
	data.Tasks = data.CreateCollection(client, "task_manager", "tasks")

	r := router.Router()
	color.Green("Server is running on port 8080")
	r.Run()
	data.CloseDB(client)
}
