package database

import (
	"errors"
	"github.com/meilisearch/meilisearch-go"
)

func ConnectMeilisearch() (client *meilisearch.Client, err error) {
	// init meilisearch client
	client = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://localhost:7700", // setup meilisearch host
		APIKey: "ThisIsMasterKey",       // setup meilisearch api key
	})

	// validate is client null or not
	if client == nil {
		return nil, errors.New("error when try to connect to meilisearch")
	}

	return
}
