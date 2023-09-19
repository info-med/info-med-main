package database

import (
	"github.com/meilisearch/meilisearch-go"
	"os"
)

var meilisearchClient *meilisearch.Client

func InitMeilisearch() {
	meilisearchClient = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: os.Getenv("MEILISEARCH_MASTER_KEY"),
	})
}
