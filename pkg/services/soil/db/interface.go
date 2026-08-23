package db

import (
	"kisaanSathi/pkg/services/soil/models"

	"github.com/jmoiron/sqlx"
)

type soilStore struct {
	db *sqlx.DB
}

func NewSoilStore(db *sqlx.DB) SoilStore {
	return &soilStore{
		db: db,
	}
}

type SoilStore interface {
	GetSoilTypes() ([]models.SoilType, error)
	GetSoilReport(userID int64) (*models.SoilReport, error)
	CreateSoilTest(req *models.CreateSoilTestRequest) error
	GetSoilRecommendation(userID int64) ([]models.SoilRecommendation, error)
}
