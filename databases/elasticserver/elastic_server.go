package elasticserver

import (
	"github.com/olivere/elastic"
	"github.com/suriyajaboon/bigquery/structs"
)

type IElasticServer interface {
	ElasticClient() (*elastic.Client, error)
	CreateIndexing(index string) (*inDiceService, error)
}

type ElasticServer struct {
	*structs.ElasticServer
}

type inDiceService struct {
	Indexing string
	State bool
	Message string
}

func ElasticNewClient(es *ElasticServer) IElasticServer {
	return &ElasticServer{ &structs.ElasticServer{es.Protocol,es.Host, es.Port}}
}

func (es *ElasticServer) ElasticClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(es.Protocol + es.Host + ":" + es.Port))
	if err != nil {
		return nil, err
	}
	return client, nil
}

