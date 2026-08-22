package controller

import (
	"kisaanSathi/pkg/services/crop/db"
	"kisaanSathi/pkg/services/crop/models"
)

type cropController struct {
	store db.CropStore
}
type CropController interface {
	GetCrops() (*models.CropListResponse, error)

	GetCrop(id int64) (*models.CropResponse, error)

	CreateCrop(req models.CreateCropRequest) error

	UpdateCrop(id int64, req models.UpdateCropRequest) error

	DeleteCrop(id int64) error

	GetCropSeason(season string) (*models.CropListResponse, error)

	GetRecommendedCrops(soil, district string) (*models.RecommendedCropResponse, error)
}

func NewCropController(store db.CropStore) CropController {
	return &cropController{
		store: store,
	}
}
