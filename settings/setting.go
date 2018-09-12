package settings

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

var environments = map[string]string{
	"production":    "settings/prod.json",
	"development": "settings/dev.json",
}

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
		log.Println("Run Application Mode " + env)
		env = "development"
	}
	LoadSettingsEnv(env)
}
func LoadSettingsEnv(env string) {
	conn, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error while reading configuration file", err)
	}
	settings = Settings{}
	jErr := json.Unmarshal(conn, &settings)
	if jErr != nil {
		fmt.Println("Error while parsing configuration file", jErr)
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

func IsModeEnvironment() bool {
	return env == "production"
}

func Cores() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		c.Next()
	}
}