package main

import (
	"github.com/gin-gonic/gin"
	"mqttx-manager/router"
)

func main() {
	engine := gin.Default()
	group := engine.Group("")

	{
		router.InitUserRouter(group)
	}

	engine.Run()
}
