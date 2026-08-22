package db

import (
	"kisaanSathi/pkg/services/metadata/models"

	"github.com/jmoiron/sqlx"
)

type metadataStore struct {
	db *sqlx.DB
}

type MetaDataStore interface {
	GetMetadata() ([]models.MetadataDB, error)
}

func NewMetadataStore(db *sqlx.DB) MetaDataStore {
	return &metadataStore{
		db: db,
	}
}