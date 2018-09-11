package routers

import (
	"github.com/gin-gonic/gin"
	"query/services"
)

func AuthenticationRoute(route *gin.Engine) *gin.Engine {
	auth := route.Group("/auth")
	//auth.Use(Authorize()) {
		auth.GET("", services.GetAuth)
		auth.GET(":id", services.GetIdAuth)
		auth.POST("", services.CreateAuth)
		auth.PUT(":id", services.UpdateAuth)
		auth.DELETE(":id", services.DeleteAuth)
	// }
	return route
}