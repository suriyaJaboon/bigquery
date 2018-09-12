package routers

import (
	"github.com/gin-gonic/gin"
	"query/services"
)

func AuthenticationRoute(route *gin.RouterGroup) {
	//auth := route.Group("/auth")
	auth := route
	//route.Use(Authorize()) {
		auth.POST("/login", services.Login)
		auth.POST("/logout", services.Logout)
	//}
}