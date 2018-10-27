package postgreserver

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/suriyajaboon/bigquery/structs"
)

type IPostgres interface {
	PostgresClient() (*gorm.DB, error)
}

type PostgresServer struct {
	*structs.PostgresServer
}

func PostgresNewClient(p *PostgresServer) IPostgres {
	return &PostgresServer{&structs.PostgresServer{p.Host, p.Port, p.User, p.Pass, p.DB}}
}

func (p *PostgresServer) PostgresClient() (*gorm.DB ,error) {
	client, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", p.Host, p.Port, p.User, p.DB, p.Pass))
	if err != nil {
		defer client.Close()
		return nil, err
	}
	return client, nil
}
