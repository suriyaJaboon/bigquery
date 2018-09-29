package settings

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
)

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

var settings Settings = Settings{}
var env = "development"

func Init() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}
	setMode(env)
	log.Println("Run Application Mode " + env)
}

func setMode(env string) {
	if env == "production" {
		settings.PrivateKeyPath = "/opt/keys/private_key"
		settings.PublicKeyPath = "/opt/keys/public_key.pub"
		settings.JWTExpirationDelta = 72
	} else {
		settings.PrivateKeyPath = "/opt/keys/private_key"
		settings.PublicKeyPath = "/opt/keys/public_key.pub"
		settings.JWTExpirationDelta = 72
	}
}


func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}

func Cores() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		c.Next()
	}
}