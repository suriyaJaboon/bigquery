package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suriyajaboon/bigquery/settings"
	"github.com/suriyajaboon/bigquery/routers"
)

func main()  {
	settings.Init()
	route := gin.Default()
	route.Use(settings.Cores())
	routers.InitRoutes(route)
	route.Run(":8090")
}

