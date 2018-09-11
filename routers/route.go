package routers

import "github.com/gin-gonic/gin"

func InitRoutes(route *gin.Engine) *gin.Engine {
	route = AuthenticationRoute(route)
	route = UserRoute(route)
	return route
}