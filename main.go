package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"log"
	"query/databases/elasticserver"
)

func main()  {
	//redis := redisserver.RedisNewClient(&redisserver.RedisServer{"localhost", "6379", "", 0})
	//clientRedis, eRedis := redis.RedisClient()
	//defer clientRedis.Close()
	//if eRedis != nil {
	//	log.Printf("%v\n", eRedis)
	//}
	//f := logs.NewCount("/home/tom/sender_zabbix_success.log")
	//getResult := clientRedis.Get(f.Count().Name())
	//fmt.Printf("%v\n", getResult)
	//
	elastic := elasticserver.ElasticNewClient(&elasticserver.ElasticServer{"http://", "localhost", "9200"})
	clientElastic, eElastic := elastic.CreateIndexing("youtube")
	if eElastic != nil {
		log.Printf("Error %v", eElastic)
	}

	fmt.Printf("Success\n %v\n %v\n %v\n", clientElastic.Indexing, clientElastic.State, clientElastic.Message)


	password := []byte("secret")
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", string(password))
	fmt.Println("Hash:    ", string(hash))

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)

	password1 := []byte("MyDarkSecret")

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password1, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, password1)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(err) // nil means it is a match
}

func HashPassword(password []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func CheckPasswordHash(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
