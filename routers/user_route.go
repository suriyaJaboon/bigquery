package routers

import (
	"github.com/gin-gonic/gin"
	"query/services"
)

func UserRoute(route *gin.RouterGroup) {
	user := route.Group("/user")
	//auth.Use(Authorize()) {
		user.GET("", services.GetUser)
		user.GET(":id", services.GetIdUser)
		user.POST("", services.CreateUser)
		user.PUT(":id", services.UpdateUser)
		user.DELETE(":id", services.DeleteUser)
	// }
}
