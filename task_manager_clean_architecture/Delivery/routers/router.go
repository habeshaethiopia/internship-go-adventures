package routers

import (
	"task/Delivery/controllers"
	infrastructure "task/Infrastructure"

	"github.com/gin-gonic/gin"
)

func Router(R *gin.Engine, secret string) {

	tc := controllers.TaskController{}
	uc := controllers.UserController{}

	// Task routes
	R.POST("/register", uc.CreateUser)
	R.POST("/login", uc.LoginUser)

	r := R.Group("/api")
	r.Use(infrastructure.AuthMiddleware(secret))
	r.GET("/tasks", tc.GetTasks)
	r.GET("/tasks/:id", tc.GetTaskByID)
	r.POST("/tasks", tc.CreateTask)
	r.PUT("/tasks/:id", tc.UpdateTask)
	r.DELETE("/tasks/:id", tc.DeleteTask)
	// User routes
	r.DELETE("/users/:id", uc.DeleteUser)
}
