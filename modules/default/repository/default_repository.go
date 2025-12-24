package repository

import (
	"context"

	"github.com/elastic/go-elasticsearch/v9"
	"gorm.io/gorm"
)

type defaultRepository struct {
	esClient *elasticsearch.Client
	db       *gorm.DB
}

func NewDefaultRepository(db *gorm.DB, es *elasticsearch.Client) DefaultRepositoryInterface {
	return &defaultRepository{
		esClient: es,
		db:       db,
	}
}

func (repo *defaultRepository) Welcome(ctx context.Context) string {
	return "Welcome to orchid-starter..."
}

func (repo *defaultRepository) WithTx(tx *gorm.DB) DefaultRepositoryInterface {
	return &defaultRepository{
		esClient: repo.esClient,
		db:       tx,
	}
}
