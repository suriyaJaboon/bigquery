package routers

import (
	"github.com/gin-gonic/gin"
	"query/services"
)

func AuthenticationRoute(route *gin.RouterGroup) {
	auth := route.Group("/auth")
	//route.Use(Authorize()) {
		auth.GET("", services.GetAuth)
		auth.GET(":id", services.GetIdAuth)
		auth.POST("", services.CreateAuth)
		auth.PUT(":id", services.UpdateAuth)
		auth.DELETE(":id", services.DeleteAuth)
	//}
}