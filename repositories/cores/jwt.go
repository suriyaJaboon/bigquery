package cores

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/suriyajaboon/bigquery/settings"
	"github.com/suriyajaboon/bigquery/structs"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"github.com/pborman/uuid"
)

type JWTAuthentication struct {
	privateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

const (
	tokenDuration = 72
	expireOffset  = 3600
)

var authBackendInstance *JWTAuthentication = nil

func NewJWTAuthenticationBackend() *JWTAuthentication {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthentication{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}
	return authBackendInstance
}

func (backend *JWTAuthentication) GenerateToken(userUUID string) (string, error){
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp" : time.Now().Add(time.Minute * time.Duration(settings.Get().JWTExpirationDelta)).Unix(),
		"iat" : time.Now().Unix(),
		"sub" : userUUID,
		"jti" : uuid.New(),
		"scopes" : "api:read",
	}
 	return token.SignedString(backend.privateKey)
}

func (backend *JWTAuthentication) Authenticate(user *structs.Users) bool {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("12345"), 10)
	testUser := structs.Users{
		UUID:     uuid.New(),
		Username: "tom",
		Password: string(hashedPassword),
	}
	return user.Username == testUser.Username && bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil
}

func (backend *JWTAuthentication) getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}

func (backend *JWTAuthentication) Logout(tokenString string, token *jwt.Token) error {
	//redisConn := redis.Connect()
	//return redisConn.SetValue(tokenString, tokenString, backend.getTokenRemainingValidity(token.Claims.(jwt.MapClaims)["exp"]))
	return nil
}

func getPrivateKey() *rsa.PrivateKey {
	pvk, err := ioutil.ReadFile(settings.Get().PrivateKeyPath)
	if err != nil {
		panic(err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pvk)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func getPublicKey() *rsa.PublicKey {
	pbk, err := ioutil.ReadFile(settings.Get().PublicKeyPath)
	if err != nil {
		panic(err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pbk)
	if err != nil {
		panic(err)
	}
	return publicKey
}