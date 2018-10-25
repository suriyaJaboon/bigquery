package elasticserver

import (
	"github.com/olivere/elastic"
)

func ElasticNewClient(es *ElasticServer) IElasticServer {
	return &ElasticServer{es.Protocol,es.Host, es.Port}
}

func (es *ElasticServer) ElasticClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(es.Protocol + es.Host + ":" + es.Port))
	if err != nil {
		return nil, err
	}
	return client, nil
}

