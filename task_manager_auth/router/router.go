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
	// User routes
	authuser:=r.Group("/auth")
	authuser.Use(middleware.AuthMiddleware())
	authuser.GET("/users", controllers.GetUsers)
	authuser.GET("/users/:id", controllers.GetUserById)
	authuser.DELETE("/user/:id", controllers.DeleteUser)
	authuser.PUT("/user/:id", controllers.UpdateUser)
	// Task routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/tasks", controllers.GetTasks)
	protected.GET("/tasks/:id", controllers.GetTaskById)
	protected.POST("/tasks", controllers.PostTask)
	protected.PUT("/tasks/:id", controllers.PutTask)
	protected.DELETE("/tasks/:id", controllers.DeleteTask)

 	return r
 }