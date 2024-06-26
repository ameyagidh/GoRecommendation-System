package db

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"graphApp/global"
)

var err error

//to initialize the driver variable
func Init() {
	global.Driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "neo4j", ""))
	if err != nil {
		panic(err)
	}

}

//to reset the database(delete all nodes and relationships)
func DeleteAll() error {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (n) DETACH DELETE n",
			map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		return nil, result.Err()
	})

	if err != nil {
		return err
	}

	return nil

}
