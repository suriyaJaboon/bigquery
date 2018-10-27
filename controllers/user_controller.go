package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/suriyajaboon/bigquery/services"
	"github.com/suriyajaboon/bigquery/structs"
	"net/http"
)

func CreateUserController(c *gin.Context) {
	userService := services.NewUserService(&services.StructUserService{&structs.Users{"d", "d", "d"}})
	gus, err := userService.GetUserService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gus)
}
