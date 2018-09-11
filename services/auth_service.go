package services

import "github.com/gin-gonic/gin"

func GetAuth(c *gin.Context) {
	c.JSON(200, gin.H{"name": "development_auth", "fullName": "bar_auth..."})
}

func GetIdAuth(c *gin.Context) {

}

func CreateAuth(c *gin.Context) {

}

func UpdateAuth(c *gin.Context) {

}

func DeleteAuth(c *gin.Context) {

}
