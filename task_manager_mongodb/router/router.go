package router
 import(
 	"github.com/gin-gonic/gin"
 	"task/controllers"
 )
 func Router() *gin.Engine {
 	r := gin.Default()
 	r.GET("/tasks", controllers.GetTask)
 	r.GET("/tasks/:id", controllers.GetTaskById)
 	r.POST("/tasks", controllers.PostTask)
 	r.PUT("/tasks/:id", controllers.PutTask)
 	r.DELETE("/tasks/:id", controllers.DeleteTask)
 	return r
 }