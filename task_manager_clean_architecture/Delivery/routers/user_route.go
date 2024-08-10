package routers

import (
	"task/Delivery/controllers"
	infrastructure "task/Infrastructure"
	repositories "task/Repositories"
	usecases "task/Usecases"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRouter(R *gin.Engine, client mongo.Database, config infrastructure.Config) {
	color.Red(config.Usercoll)
	ur := repositories.NewUserRepository(client, config.Usercoll)
	jwtService := infrastructure.NewJWTService(config.Jwt_secret)
	uu := usecases.NewUserUsecase(ur, *jwtService)
	uc := &controllers.UserController{
		UserUsecase: uu,
	}
	// Task routes
	R.POST("/register", uc.CreateUser)
	R.POST("/login", uc.LoginUser)
	r := R.Group("/api")
	r.Use(infrastructure.AuthMiddleware([]byte(config.Jwt_secret)))
	r.GET("/user/:id", uc.GetUserByID)
	r.DELETE("/user/:id", uc.DeleteUser)
	r.GET("/users", uc.GetUsers)

}
