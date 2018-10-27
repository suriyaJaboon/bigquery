package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/suriyajaboon/bigquery/controllers"
)

func UserRoute(route *gin.RouterGroup) {
	user := route.Group("/user")
	user.GET("", controllers.CreateUserController)
	//user.GET(":id", commons.GetIdUser)
	//user.POST("", commons.CreateUser)
	//user.PUT(":id", commons.UpdateUser)
	//user.DELETE(":id", commons.DeleteUser)
}
