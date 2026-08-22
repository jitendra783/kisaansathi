package db

import (
	"kisaanSathi/pkg/services/crop/models"

	"github.com/jmoiron/sqlx"
)

type cropStore struct {
	db *sqlx.DB
}

type CropStore interface {
	GetCrops() ([]models.Crop, error)

	GetCrop(id int64) (*models.Crop, error)

	CreateCrop(req models.CreateCropRequest) error

	UpdateCrop(id int64, req models.UpdateCropRequest) error

	DeleteCrop(id int64) error

	GetCropSeason(season string) ([]models.Crop, error)

	GetRecommendedCrops(soil, district string) ([]models.RecommendedCrop, error)
}

func NewCropStore(db *sqlx.DB) CropStore {
	return &cropStore{db: db}
}
