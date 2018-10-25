package postgreserver

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type PostgressServer struct {
	Host string
	Port string
	Username string
	Password string
	Database string
}

type IPostgres interface {
	PostgresClient() (*gorm.DB, error)
}

func PostgresNewClient(p *PostgressServer) IPostgres {
	return &PostgressServer{p.Host, p.Port, p.Username, p.Password, p.Database}
}


func (p *PostgressServer) PostgresClient() (*gorm.DB ,error) {
	client, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", p.Host, p.Port, p.Username, p.Database, p.Password))
	if err != nil {
		defer client.Close()
		return nil, err
	}
	return client, nil
}
