package db

import (
	"kisaanSathi/pkg/services/mandi/models"

	"github.com/jmoiron/sqlx"
)

type mandiStore struct {
	db *sqlx.DB
}

type MandiStore interface {
	GetMandiBhav() ([]models.MandiPrice, error)
	GetMandiPrices() ([]models.MandiPrice, error)
	GetCropPrices(crop string) ([]models.MandiPrice, error)
	GetStatePrices(state string) ([]models.MandiPrice, error)
	GetDistrictPrices(district string) ([]models.MandiPrice, error)
	GetTrendingPrices() ([]models.TrendingPrice, error)
	GetPriceComparison(crop, market string) ([]models.PriceHistory, error)
}

func NewMandiStore(db *sqlx.DB) MandiStore {
	return &mandiStore{db: db}
}
