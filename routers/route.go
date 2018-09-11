package routers

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	route := gin.Default()
	return route
}