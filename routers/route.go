package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/suriyajaboon/bigquery/controllers"
	"github.com/suriyajaboon/bigquery/commons/cores"
)

func InitRoutes(ge *gin.Engine) {
	ge.POST("/api/login", controllers.Login)
	route := ge.Group("/v1", cores.JWTAuthMiddleware())
	AuthenticationRoute(route)
	UserRoute(route)
}