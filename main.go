package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/suriyajaboon/bigquery/routers"
	"github.com/suriyajaboon/bigquery/settings"
	"io"
	"os"
)

func main()  {
	settings.Init()
	//Log file

	conf, err := settings.GetConfiguration()
	if err != nil {
		fmt.Println(err)
	}
	f, _ := os.Create(conf.Logs.Application)
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
	route.Run( "localhost:" + settings.GetPort())
}

