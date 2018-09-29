package cores

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"fmt"
	"net/http"
	"github.com/suriyajaboon/bigquery/structs"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := GetToken(c)
		token, err := ValidateToken(tokenString)
		if err == nil && token.Valid && NewJWTAuthenticationBackend().IsInBlacklist(token) {
			c.Set("uuid", token.Claims.(jwt.MapClaims)["jti"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, &structs.TokenAuthentication{http.StatusUnauthorized, "", "Unauthorized"})
			c.Abort()
		}
	}
}

func GetToken(c *gin.Context) string {
	tokenString := c.DefaultQuery("token", "")
	if tokenString == "" {
		tokenString = c.Request.Header.Get("Authorization")
		if s := strings.Split(tokenString, " "); len(s) == 2 {
			tokenString = s[1]
		}
	}
	if len(tokenString) == 0 {
		return ""
	}
	return tokenString
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return NewJWTAuthenticationBackend().PublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (backend *JWTAuthentication) IsInBlacklist(token *jwt.Token) bool {
	//redisConn := redis.Connect()
	//redisToken, _ := redisConn.GetValue(token)
	return true
}