package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/suriyajaboon/bigquery/controllers"
)

func AuthenticationRoute(route *gin.RouterGroup) {
	//auth := route.Group("/auth")
	//route.POST("/login", controllers.Login)
	route.GET("/logout", controllers.Logout)
}