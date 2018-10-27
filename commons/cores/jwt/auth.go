package jwt
//
//import (
//	"github.com/dgrijalva/jwt-go"
//	"crypto/rsa"
//	"time"
//	"query/settings"
//	"os"
//	"bufio"
//	"encoding/pem"
//	"crypto/x509"
//	"golang.org/x/crypto/bcrypt"
//	"query/structs"
//	"github.com/pborman/uuid"
//	"fmt"
//)
//
//type JWTAuthentication struct {
//	privateKey *rsa.PrivateKey
//	PublicKey *rsa.PublicKey
//}
//
//const (
//	tokenDuration = 72
//	expireOffset  = 3600
//)
//
//var authBackendInstance *JWTAuthentication = nil
//
//func NewJWTAuthenticationBackend() *JWTAuthentication {
//	if authBackendInstance == nil {
//		authBackendInstance = &JWTAuthentication{
//			privateKey: getPrivateKey(),
//			PublicKey:  getPublicKey(),
//		}
//	}
//	return authBackendInstance
//}
//
//func (backend *JWTAuthentication) GenerateToken(userUUID string) (string, error){
//	token := jwt.New(jwt.SigningMethodRS512)
//	token.Claims = jwt.MapClaims{
//		"ExpiresAt" : time.Now().Add(time.Hour * time.Duration(settings.Get().JWTExpirationDelta)).Unix(),
//		"IssuedAt" : time.Now().Unix(),
//		"Issuer" : userUUID,
//	}
//	tokenString, err := token.SignedString(backend.privateKey)
//	if err != nil {
//		panic(err)
//		return "", err
//	}
//	fmt.Println(tokenString)
//	return tokenString, nil
//}
//
//func (backend *JWTAuthentication) Authenticate(user *structs.Users) bool {
//	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("12345"), 10)
//
//	testUser := structs.Users{
//		UUID:     uuid.New(),
//		Username: "tom",
//		Password: string(hashedPassword),
//	}
//	return user.Username == testUser.Username && bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil
//}
//
//func (backend *JWTAuthentication) getTokenRemainingValidity(timestamp interface{}) int {
//	if validity, ok := timestamp.(float64); ok {
//		tm := time.Unix(int64(validity), 0)
//		remainer := tm.Sub(time.Now())
//		if remainer > 0 {
//			return int(remainer.Seconds() + expireOffset)
//		}
//	}
//	return expireOffset
//}
//
//func (backend *JWTAuthentication) IsInBlacklist(token string) bool {
//	//redisConn := redis.Connect()
//	//redisToken, _ := redisConn.GetValue(token)
//
//	//if redisToken == nil {
//	//	return false
//	//}
//	return true
//}
//
//func (backend *JWTAuthentication) Logout(tokenString string, token *jwt.Token) error {
//	//redisConn := redis.Connect()
//	//return redisConn.SetValue(tokenString, tokenString, backend.getTokenRemainingValidity(token.Claims.(jwt.MapClaims)["exp"]))
//	return nil
//}
//
//func getPrivateKey() *rsa.PrivateKey {
//	privateKeyFile, err := os.Open(settings.Get().PrivateKeyPath)
//	if err != nil {
//		panic(err)
//	}
//
//	pemfileinfo, _ := privateKeyFile.Stat()
//	pembytes := make([]byte, pemfileinfo.Size())
//
//	buffer := bufio.NewReader(privateKeyFile)
//	_, err = buffer.Read(pembytes)
//	data, _ := pem.Decode([]byte(pembytes))
//	privateKeyFile.Close()
//
//	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
//	if err != nil {
//		panic(err)
//	}
//	return privateKeyImported
//}
//
//func getPublicKey() *rsa.PublicKey {
//	publicKeyFile, err := os.Open(settings.Get().PublicKeyPath)
//	if err != nil {
//		panic(err)
//	}
//
//	pemfileinfo, _ := publicKeyFile.Stat()
//	pembytes := make([]byte, pemfileinfo.Size())
//
//	buffer := bufio.NewReader(publicKeyFile)
//	_, err = buffer.Read(pembytes)
//	data, _ := pem.Decode([]byte(pembytes))
//	publicKeyFile.Close()
//
//	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
//	if err != nil {
//		panic(err)
//	}
//	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)
//	if !ok {
//		panic(err)
//	}
//	return rsaPub
//}