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
var port = "8090"

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}
	setMode(env)
	log.Printf("Application Mode %s Run localhost:%s", env, GetPort())
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

func GetPort() string {
	port = os.Getenv("GO_PORT")
	if port == "" {
		port = "8090"
	}
	return port
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