package controller

import (
	"kisaanSathi/pkg/services/metadata/db"
	"kisaanSathi/pkg/services/metadata/models"
)

type controller struct {
	store db.MetaDataStore
}
type MetadataController interface {
	GetMetadata() (*models.MetadataResponse, error)
}

func NewMetadataController(store db.MetaDataStore) MetadataController {
	return &controller{
		store: store,
	}
}
