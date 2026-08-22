package controller

import (
	"kisaanSathi/pkg/services/contact/db"
	"kisaanSathi/pkg/services/contact/models"
)

type contactController struct {
	store db.ContactStore
}
type ContactController interface {
	ContactUs(req *models.ContactRequest) (*models.ContactResponse, error)
	GetContactInfo() (*models.ContactInfoResponse, error)
}

func NewContactController(store db.ContactStore) ContactController {
	return &contactController{
		store: store,
	}
}
