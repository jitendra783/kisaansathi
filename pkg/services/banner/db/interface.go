package db

import (
	"context"

	"kisaanSathi/pkg/services/banner/models"

	"github.com/jmoiron/sqlx"
)

type dbObject struct {
	db *sqlx.DB
}

type BannerStore interface {

	// Banner
	GetBanners(ctx context.Context) ([]*models.Banner, error)

	GetBannerByID(ctx context.Context, id int64) (*models.Banner, error)

	CreateBanner(ctx context.Context, banner *models.Banner) error

	UpdateBanner(ctx context.Context, banner *models.Banner) error

	DeleteBanner(ctx context.Context, id int64) error

	GetActiveBanners(ctx context.Context) ([]*models.Banner, error)
}

func NewDBObject(db *sqlx.DB) BannerStore {
	return &dbObject{
		db: db,
	}
}
