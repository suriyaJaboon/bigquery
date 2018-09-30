package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suriyajaboon/bigquery/settings"
	"github.com/suriyajaboon/bigquery/routers"
	"os"
	"io"
)

func main()  {
	settings.Init()

	// Log file
	f, _ := os.Create("bigQuery.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Configuration Production Mode
	gin.SetMode(gin.ReleaseMode)
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

	// Development Mode
	//route := gin.Default()

	route.Use(settings.Cores())
	routers.InitRoutes(route)
	route.Run( "localhost:8090")
}

