package controller

import (
	"context"
	"errors"

	"kisaanSathi/pkg/services/expert/db"
	"kisaanSathi/pkg/services/expert/models"
)

type expertController struct {
	store db.ExpertStore
}

type ExpertController interface {
	GetExperts(ctx context.Context) (*models.GetExpertsResponse, error)
	GetExpert(ctx context.Context, id int64) (*models.GetExpertResponse, error)
	GetExpertsByCategory(ctx context.Context, category string) (*models.GetExpertsResponse, error)
	BookConsultation(ctx context.Context, req *models.BookConsultationRequest) (*models.BookConsultationResponse, error)
	GetConsultationHistory(ctx context.Context, userID int64) (*models.ConsultationHistoryResponse, error)
}

func NewExpertController(store db.ExpertStore) ExpertController {
	return &expertController{
		store: store,
	}
}

// GetExperts returns all experts.
func (c *expertController) GetExperts(ctx context.Context) (*models.GetExpertsResponse, error) {

	result, err := c.store.GetExperts(ctx)
	if err != nil {
		return nil, err
	}

	return &models.GetExpertsResponse{
		Data: result,
	}, nil
}

// GetExpert returns expert details by ID.
func (c *expertController) GetExpert(ctx context.Context, id int64) (*models.GetExpertResponse, error) {

	result, err := c.store.GetExpert(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("expert not found")
	}

	return &models.GetExpertResponse{
		Data: result,
	}, nil
}

// GetExpertsByCategory returns experts by specialization.
func (c *expertController) GetExpertsByCategory(ctx context.Context, category string) (*models.GetExpertsResponse, error) {

	result, err := c.store.GetExpertsByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return &models.GetExpertsResponse{
		Data: result,
	}, nil
}

// BookConsultation books a consultation with an expert.
func (c *expertController) BookConsultation(
	ctx context.Context,
	req *models.BookConsultationRequest,
) (*models.BookConsultationResponse, error) {

	err := c.store.BookConsultation(ctx, req)
	if err != nil {
		return nil, err
	}

	return &models.BookConsultationResponse{
		Message: "Consultation booked successfully.",
	}, nil
}

// GetConsultationHistory returns user's consultation history.
func (c *expertController) GetConsultationHistory(
	ctx context.Context,
	userID int64,
) (*models.ConsultationHistoryResponse, error) {

	result, err := c.store.GetConsultationHistory(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &models.ConsultationHistoryResponse{
		Data: result,
	}, nil
}