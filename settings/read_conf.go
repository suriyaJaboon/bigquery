package settings

import (
	"fmt"
	"encoding/json"
	"github.com/suriyajaboon/bigquery/structs"
	"os"
)

func LoadFile(path string) (*structs.ObjectType, error) {
	var configuration *structs.ObjectType
	conf, err := os.Open(path)
	defer conf.Close()
	if err != nil {
		fmt.Println(err)
	}
	jsonParser := json.NewDecoder(conf)
	jsonParser.Decode(&configuration)
	return configuration, err
}