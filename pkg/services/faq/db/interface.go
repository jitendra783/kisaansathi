package db

import (
	"kisaanSathi/pkg/services/faq/models"

	"github.com/jmoiron/sqlx"
)

type faqStore struct {
	db *sqlx.DB
}
type FAQStore interface {
	GetFAQs() ([]models.FAQ, error)
	GetFAQByID(id int64) (*models.FAQ, error)
	GetFAQsByCategory(category string) ([]models.FAQ, error)
}

func NewFAQStore(db *sqlx.DB) FAQStore {
	return &faqStore{
		db: db,
	}
}
