package main

import (
	"fmt"
	"task/Delivery/routers"
	infrastructure "task/Infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	server:= gin.Default()
	config, err := infrastructure.LoadEnv()
	if err != nil {
		fmt.Print("Error in env.load" )
	}
	fmt.Print(config)

	routers.Router(server,config.Jwt_secret)
	server.Run(fmt.Sprintf(":%d", config.Port))
}
