package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/suriyajaboon/bigquery/structs"
	"github.com/suriyajaboon/bigquery/repositories/services"
	"net/http"
	"fmt"
)

func Login(c *gin.Context) {
	var user *structs.Users
	c.BindJSON(&user)
	responseStatus, object := services.Login(user)
	c.JSON(responseStatus, object)
}

func Logout(c *gin.Context) {
	loggedInInterface, exists := c.Get("uuid")
	fmt.Println(loggedInInterface)
	fmt.Println(exists)
	c.JSON(http.StatusOK, services.Logout())
}