package db

import (
	"kisaanSathi/pkg/services/schemes/models"

	"github.com/jmoiron/sqlx"
)

type schemeStore struct {
	db *sqlx.DB
}

type SchemeStore interface {
	GetSchemes(page, limit int) ([]models.Scheme, int, error)
	GetScheme(id int64) (*models.Scheme, error)
	GetEligibleSchemes(state, category, farmerType string) ([]models.Scheme, error)
	GetStateSchemes(state string) ([]models.Scheme, error)
}

func NewSchemeStore(db *sqlx.DB) SchemeStore {
	return &schemeStore{
		db: db,
	}
}
