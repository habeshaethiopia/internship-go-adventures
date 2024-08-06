package router

import (
	"task/controllers"
	"task/middleware"

	"github.com/gin-gonic/gin"
)
 func Router() *gin.Engine {
 	r := gin.Default()
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

 	r.GET("/tasks",middleware.AuthMiddleware(), controllers.GetTask)
 	r.GET("/tasks/:id", controllers.GetTaskById)
 	r.POST("/tasks", controllers.PostTask)
 	r.PUT("/tasks/:id", controllers.PutTask)
 	r.DELETE("/tasks/:id", controllers.DeleteTask)
 	return r
 }