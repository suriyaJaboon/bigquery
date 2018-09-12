package services

import (
	"github.com/gin-gonic/gin"
	"query/cores/authentication"
	"query/models"
	"net/http"
)

func Login(c *gin.Context) {
	var user *models.Users
	c.Bind(&user)
	auThen := authentication.NewJWTAuthenticationBackend()
	if auThen.Authenticate(user) {
		toKen, err := auThen.GenerateToken(user.UUID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Server Error")
		} else {
			c.JSON(http.StatusOK, models.TokenAuthentication{toKen})
			//c.Abort()
		}
	}
	//c.JSON(http.StatusUnauthorized, "Unauthorized")
}

func Logout(c *gin.Context) {

}

