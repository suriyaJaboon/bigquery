package elasticserver

import "github.com/olivere/elastic"

type IElasticServer interface {
	ElasticClient() (*elastic.Client, error)
	CreateIndexing(index string) (*inDiceService, error)
}

type ElasticServer struct {
	Protocol string
	Host string
	Port string
}

type inDiceService struct {
	Indexing string
	State bool
	Message string
}