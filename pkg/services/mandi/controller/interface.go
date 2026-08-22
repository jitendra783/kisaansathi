package controller

import (
	"kisaanSathi/pkg/services/mandi/db"
	"kisaanSathi/pkg/services/mandi/models"
)

type mandiController struct {
	store db.MandiStore
}
type MandiController interface {
	GetMandiBhav() ([]models.MandiPriceResponse, error)
	GetMandiPrices() ([]models.MandiPriceResponse, error)
	GetCropPrices(crop string) ([]models.MandiPriceResponse, error)
	GetStatePrices(state string) ([]models.MandiPriceResponse, error)
	GetDistrictPrices(district string) ([]models.MandiPriceResponse, error)
	GetTrendingPrices() ([]models.TrendingPriceResponse, error)
		GetPriceComparison(crop, market string) (*models.PriceComparisonResponse, error)

}

func NewMandiController(store db.MandiStore) MandiController {
	return &mandiController{
		store: store,
	}
}
