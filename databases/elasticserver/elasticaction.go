package elasticserver

import (
	"context"
	"github.com/olivere/elastic"
	"fmt"
)

func CheckingIndexing(client *elastic.Client, indexName string) (bool, error) {
	exists, err := client.IndexExists(indexName).Do(context.Background())
	if err != nil {
		return false, err
	}
	if exists {
		return exists, nil
	} else {
		return exists, nil
	}
}

func (es *ElasticServer) CreateIndexing(index string) (*inDiceService, error) {
	cCon, eCon := es.ElasticClient()
	if eCon != nil {
		return &inDiceService{index, false, "Failed To Connection ElasticSearch"}, eCon
	}

	inDice, errDice := CheckingIndexing(cCon, index)
	if errDice != nil {
		return &inDiceService{index, false, "Failed To Create Indexing"}, errDice
	}
	if inDice {
		return &inDiceService{index, false, "Indexing ถูกสร้างไว้แล้ว"}, errDice
	} else {
		cInDice, cErr := cCon.CreateIndex(index).Do(context.Background())
		if (cErr != nil) && (cInDice.Index != index) {
			fmt.Printf("%v, %v", cErr, cInDice.Index)
			return &inDiceService{index, false, "Create Indexing Failed"}, cErr
		}
		return &inDiceService{cInDice.Index, true, "Create Indexing Successfully"}, nil
	}
}
