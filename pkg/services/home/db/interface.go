package db

import (
	"context"
	"kisaanSathi/pkg/services/home/models"

	"github.com/jmoiron/sqlx"
)

type homeStore struct {
	db *sqlx.DB
}

type HomeStore interface {
	GetDashboard(ctx context.Context) (*models.DashboardResponse, error)
	GetHomeData(ctx context.Context) (*models.HomeResponse, error)
}

func NewHomeStore(db *sqlx.DB) HomeStore {
	return &homeStore{
		db: db,
	}
}
