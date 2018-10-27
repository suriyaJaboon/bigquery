package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/suriyajaboon/bigquery/databases/postgreserver"
	"github.com/suriyajaboon/bigquery/services"
	"github.com/suriyajaboon/bigquery/structs"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	var user *structs.Users
	c.BindJSON(&user)
	responseStatus, object := services.Login(user)
	con,err := postgreserver.PostgresNewClient(&postgreserver.PostgresServer{ &structs.PostgresServer{"localhost", "5432", "postgres", "admin", "lab"}}).PostgresClient()
	if err != nil {
		log.Panicln(err)
	}
	services.NewLab(con).Create()
	defer con.Close()
	c.JSON(responseStatus, object)
}

func Logout(c *gin.Context) {
	loggedInInterface, exists := c.Get("uuid")
	fmt.Println(loggedInInterface)
	fmt.Println(exists)
	c.JSON(http.StatusOK, services.Logout())
}