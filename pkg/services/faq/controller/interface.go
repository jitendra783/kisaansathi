package controller

import (
	"kisaanSathi/pkg/services/faq/db"
	"kisaanSathi/pkg/services/faq/models"
)

type faqController struct {
	store db.FAQStore
}
type FAQController interface {
	GetFAQs() (*models.FAQResponse, error)
	GetFAQByID(id int64) (*models.FAQDetailResponse, error)
	GetFAQsByCategory(category string) (*models.FAQResponse, error)
}

func NewFAQController(store db.FAQStore) FAQController {
	return &faqController{
		store: store,
	}
}
