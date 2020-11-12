package router

import (
	"github.com/gin-gonic/gin"
	"mqttx-manager/controller"
	"mqttx-manager/middleware"
)

func InitUserRouter(group *gin.RouterGroup) {
	r := group.Group("/user")
	r.Use(middleware.AuthFilter)
	r.POST("/signIn", controller.SignIn)

	gin.Recovery()
}
