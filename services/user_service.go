package services

import "github.com/gin-gonic/gin"

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{"name": "development_user", "fullName": "bar_user..."})
}

func GetIdUser(c *gin.Context) {

}

func CreateUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
