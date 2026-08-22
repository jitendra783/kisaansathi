package controller

import (
	"kisaanSathi/pkg/services/schemes/db"
	"kisaanSathi/pkg/services/schemes/models"
)

type schemeController struct {
	store db.SchemeStore
}

type SchemeController interface {
	GetSchemes(page, limit int) (*models.SchemeResponse, error)
	GetScheme(id int64) (*models.SchemeDetailResponse, error)
	GetEligibleSchemes(state, category, farmerType string) (*models.SchemeResponse, error)
	GetStateSchemes(state string) (*models.SchemeResponse, error)
}

func NewSchemeController(store db.SchemeStore) SchemeController {
	return &schemeController{
		store: store,
	}
}
