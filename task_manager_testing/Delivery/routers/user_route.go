package routers

import (
	"task/Delivery/controllers"
	infrastructure "task/Infrastructure"
	repositories "task/Repositories"
	usecases "task/Usecases"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"task/mongo"
)

func UserRouter(R *gin.Engine, client mongo.Database, config infrastructure.Config) {
	color.Red(config.Usercoll)
	ur := repositories.NewUserRepository(client, config.Usercoll)
	jwtService := infrastructure.NewJWTService(config.Jwt_secret)
	uu := usecases.NewUserUsecase(ur, *jwtService)
	uc := controllers.NewUserController(uu)
	
	// Task routes
	R.POST("/register", uc.CreateUser)
	R.POST("/login", uc.LoginUser)
	r := R.Group("/api")
	r.Use(infrastructure.AuthMiddleware([]byte(config.Jwt_secret)))
	r.GET("/users/:id", uc.GetUserByID)
	r.DELETE("/users/:id", uc.DeleteUser)
	r.GET("/users", uc.GetUsers)
	r.PUT("/users/:id", uc.UpdateUser)

}
