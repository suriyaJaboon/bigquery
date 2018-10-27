package settings

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/suriyajaboon/bigquery/structs"
	"log"
	"os"
	"path/filepath"
)

var settings structs.SettingMode = structs.SettingMode{}
var env string
var port string

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}
	setMode()
	log.Printf("Application Mode %s Run localhost:%s", env, GetPort())
}

func setMode() {
	conf, err := GetConfiguration()
	if err != nil {
		fmt.Println(err)
	}
	settings.PrivateKeyPath = conf.SettingMode.PrivateKeyPath
	settings.PublicKeyPath = conf.SettingMode.PublicKeyPath
	settings.JWTExpirationDelta = conf.SettingMode.JWTExpirationDelta
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

func Get() structs.SettingMode {
	if &settings == nil {
		Init()
	}
	return settings
}

func GetConfiguration() (*structs.ObjectType, error) {
	var pathConf string
	pathConf = "/Users/suriya/go/src/github.com/suriyajaboon/bigquery/" + GetEnvironment() + ".json"
	if GetEnvironment() == "production" {
		path, err := os.Executable()
		if err != nil {
			log.Printf("%v", err)
		}
		pathConf = filepath.Dir(path) + "/" + GetEnvironment() + ".json"
	}
	return LoadFile(pathConf)
}

func Cores() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		c.Next()
	}
}