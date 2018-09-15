package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/suriyajaboon/bigquery/controllers"
)

func UserRoute(route *gin.RouterGroup) {
	user := route.Group("/user")
	user.GET("", controllers.CreateUserController)
	//user.GET(":id", repositories.GetIdUser)
	//user.POST("", repositories.CreateUser)
	//user.PUT(":id", repositories.UpdateUser)
	//user.DELETE(":id", repositories.DeleteUser)
}
