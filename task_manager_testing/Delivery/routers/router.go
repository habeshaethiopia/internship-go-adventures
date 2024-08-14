package routers

import (
	infrastructure "task/Infrastructure"

	"task/mongo"

	"github.com/gin-gonic/gin"
)

func Router(R *gin.Engine, env infrastructure.Config, client mongo.Database) {

	UserRouter(R, client, env)
	TaskRouter(R, client, env)

}
