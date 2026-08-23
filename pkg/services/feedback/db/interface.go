package db

import (
	"kisaanSathi/pkg/services/feedback/models"

	"github.com/jmoiron/sqlx"
)

type feedbackStore struct {
	db *sqlx.DB
}
type FeedbackStore interface {
	CreateFeedback(req *models.FeedbackRequest) error
	GetFeedbacks() ([]models.Feedback, error)
}

func NewFeedbackStore(db *sqlx.DB) FeedbackStore {
	return &feedbackStore{
		db: db,
	}
}
