package main

import (
	"github.com/gin-gonic/gin"
	"query/settings"
	"query/routers"
)

func main()  {
	settings.Init()
	route := gin.Default()
	route.Use(settings.Cores())
	v1 := route.Group("/v1")
	routers.InitRoutes(v1)
	route.Run(":8090")
}

