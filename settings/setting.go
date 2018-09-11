package settings

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
)

var env = "development"

func Init() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		log.Println("Run Application Mode Development")
		env = "development"
	}
}

func GetEnvironment() string {
	return env
}

func Cores() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		c.Next()
	}
}