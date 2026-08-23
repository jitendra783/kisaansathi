package db

import (
	"kisaanSathi/pkg/services/contact/models"

	"github.com/jmoiron/sqlx"
)

type contactStore struct {
	db *sqlx.DB
}
type ContactStore interface {
	SaveContact(req *models.ContactRequest) error
	GetContactInfo() (*models.ContactInfo, error)
}

func NewContactStore(db *sqlx.DB) ContactStore {
	return &contactStore{
		db: db,
	}
}
