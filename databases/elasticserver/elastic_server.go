package elasticserver

import (
	"github.com/olivere/elastic"
)

func ElasticNewClient(es *ElasticServer) IElasticServer {
	return &ElasticServer{es.Protocol,es.Host, es.Port}
}

func (es *ElasticServer) ElasticClient() (*elastic.Client, error) {
	client, err := es.connectToElastic()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (es *ElasticServer) connectToElastic() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(es.Protocol + es.Host + ":" + es.Port))
	if err != nil {
		return nil, err
	}
	return client, nil
}
