package db

import (
	"kisaanSathi/pkg/services/feedback/models"
)

func (s *feedbackStore) CreateFeedback(req *models.FeedbackRequest) error {

	query := `
		INSERT INTO feedback
		(
			name,
			phone,
			email,
			rating,
			category,
			message,
			app_version,
			created_at
		)
		VALUES
		(
			:name,
			:phone,
			:email,
			:rating,
			:category,
			:message,
			:app_version,
			NOW()
		)
	`

	_, err := s.db.NamedExec(query, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *feedbackStore) GetFeedbacks() ([]models.Feedback, error) {

	var feedbacks []models.Feedback

	query := `
		SELECT
			id,
			name,
			phone,
			email,
			rating,
			category,
			message,
			app_version,
			created_at
		FROM feedback
		ORDER BY created_at DESC
	`

	err := s.db.Select(&feedbacks, query)
	if err != nil {
		return nil, err
	}

	return feedbacks, nil
}
