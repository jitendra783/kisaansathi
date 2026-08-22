package controller

import "kisaanSathi/pkg/services/feedback/models"

func (c *feedbackController) CreateFeedback(req *models.FeedbackRequest) (*models.SuccessResponse, error) {

	if err := c.store.CreateFeedback(req); err != nil {
		return nil, err
	}

	return &models.SuccessResponse{
		Message: "Feedback submitted successfully",
	}, nil
}

func (c *feedbackController) GetFeedbacks() (*models.FeedbackListResponse, error) {

	result, err := c.store.GetFeedbacks()
	if err != nil {
		return nil, err
	}

	return &models.FeedbackListResponse{
		Data: result,
	}, nil
}
