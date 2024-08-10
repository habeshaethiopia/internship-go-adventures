package main

import (
	"context"
	"fmt"

	"task/Delivery/routers"
	infrastructure "task/Infrastructure"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	config, err := infrastructure.LoadEnv()
	if err != nil {
		fmt.Print("Error in env.load")
	}
	fmt.Print(config)
	DB, client, err := infrastructure.ConnectDB(config.DatabaseUrl, config.Dbname)
	
	if err != nil {
		fmt.Print("Error in connectDB")
	}
	coll:=DB.Collection("users")
	coll.InsertOne(context.Background(), gin.H{"name": "pi", "value": 3.14159})
	color.Red(coll.Name())
	defer infrastructure.CloseDB(client)
	routers.Router(server, *config, DB)
	server.Run(fmt.Sprintf(":%d", config.Port))

	
}
