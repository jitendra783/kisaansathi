package controller

import (
	"kisaanSathi/pkg/services/feedback/db"
	"kisaanSathi/pkg/services/feedback/models"
)

type feedbackController struct {
	store db.FeedbackStore
}
type FeedbackController interface {
	CreateFeedback(req *models.FeedbackRequest) (*models.SuccessResponse, error)
	GetFeedbacks() (*models.FeedbackListResponse, error)
}

func NewFeedbackController(store db.FeedbackStore) FeedbackController {
	return &feedbackController{
		store: store,
	}
}
