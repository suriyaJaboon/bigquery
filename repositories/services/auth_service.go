package services

import (
	"github.com/suriyajaboon/bigquery/repositories/cores"
	"github.com/suriyajaboon/bigquery/structs"
	"net/http"
	"github.com/gin-gonic/gin/json"
)

func Login(user *structs.Users) (int, *structs.TokenAuthentication) {
	authBackend := cores.NewJWTAuthenticationBackend()
	if authBackend.Authenticate(user) {
		toKen, err := authBackend.GenerateToken(user.UUID)
		if err != nil {
			return http.StatusInternalServerError, &structs.TokenAuthentication{http.StatusInternalServerError, "", "Server Error"}
		} else {
			return http.StatusOK, &structs.TokenAuthentication{http.StatusOK, toKen, "Successfully"}
		}
	}
	return  http.StatusUnauthorized, &structs.TokenAuthentication{http.StatusUnauthorized, "", "Unauthorized"}
}

func RefreshToken(requestUser *structs.Users) []byte {
	token, err := cores.NewJWTAuthenticationBackend().GenerateToken(requestUser.UUID)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(structs.TokenAuthentication{http.StatusOK, token, "Successfully"})
	if err != nil {
		panic(err)
	}
	return response
}
func Logout() bool {
	return true
}

