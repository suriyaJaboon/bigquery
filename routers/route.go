package routers

import "github.com/gin-gonic/gin"

func InitRoutes(route *gin.RouterGroup) {
	AuthenticationRoute(route)
	UserRoute(route)
}