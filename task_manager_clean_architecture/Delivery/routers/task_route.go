package routers

import (
	"task/Delivery/controllers"
	infrastructure "task/Infrastructure"
	repositories "task/Repositories"
	usecases "task/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func TaskRouter(R *gin.Engine, client mongo.Database, config infrastructure.Config) {

	tr := repositories.NewTaskRepository(client, config.Taskcoll)
	tu := usecases.NewTaskUsecase(tr)
	tc := &controllers.TaskController{
		TaskUsecase: tu,
	}
	// Task routes
	r := R.Group("/api")
	r.Use(infrastructure.AuthMiddleware([]byte(config.Jwt_secret)))
	r.GET("/tasks", tc.GetTasks)
	r.GET("/tasks/:id", tc.GetTaskByID)
	r.POST("/task", tc.CreateTask)
	r.PUT("/tasks/:id", tc.UpdateTask)
	r.DELETE("/tasks/:id", tc.DeleteTask)

}
