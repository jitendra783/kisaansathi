package controller

import (
	"context"

	"kisaanSathi/pkg/services/banner/db"
	"kisaanSathi/pkg/services/banner/models"
)

type controller struct {
	store db.BannerStore
}

type BannerController interface {

	// Banner
	GetBanners(ctx context.Context) ([]*models.BannerResponse, error)

	GetBannerByID(ctx context.Context, id int64) (*models.BannerResponse, error)

	CreateBanner(ctx context.Context, req *models.CreateBannerRequest) error

	UpdateBanner(ctx context.Context, req *models.UpdateBannerRequest) error

	DeleteBanner(ctx context.Context, id int64) error

	GetActiveBanners(ctx context.Context) ([]*models.BannerResponse, error)
}

func NewBannerController(store db.BannerStore) BannerController {
	return &controller{
		store: store,
	}
}