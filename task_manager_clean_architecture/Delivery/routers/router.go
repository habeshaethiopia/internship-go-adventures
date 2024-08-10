package routers

import (
	infrastructure "task/Infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(R *gin.Engine, env infrastructure.Config, client *mongo.Database) {

	UserRouter(R, *client, env)
	TaskRouter(R, *client, env)

}
