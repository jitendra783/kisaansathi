package db

import (
	"context"
	"database/sql"

	"kisaanSathi/pkg/services/expert/models"

	"github.com/jmoiron/sqlx"
)

type expertStore struct {
	db *sqlx.DB
}

type ExpertStore interface {
	GetExperts(ctx context.Context) ([]models.Expert, error)
	GetExpert(ctx context.Context, id int64) (*models.Expert, error)
	GetExpertsByCategory(ctx context.Context, category string) ([]models.Expert, error)
	BookConsultation(ctx context.Context, req *models.BookConsultationRequest) error
	GetConsultationHistory(ctx context.Context, userID int64) ([]models.ConsultationHistory, error)
}

func NewExpertStore(db *sqlx.DB) ExpertStore {
	return &expertStore{
		db: db,
	}
}

// GetExperts returns all available experts.
func (s *expertStore) GetExperts(ctx context.Context) ([]models.Expert, error) {

	var experts []models.Expert

	query := `
		SELECT
			id,
			name,
			specialization,
			experience,
			language,
			rating,
			bio,
			consultation_fee,
			phone,
			profile_image,
			available
		FROM experts
		ORDER BY rating DESC, experience DESC;
	`

	err := s.db.SelectContext(ctx, &experts, query)
	if err != nil {
		return nil, err
	}

	return experts, nil
}

// GetExpert returns expert details by id.
func (s *expertStore) GetExpert(ctx context.Context, id int64) (*models.Expert, error) {

	var expert models.Expert

	query := `
		SELECT
			id,
			name,
			specialization,
			experience,
			language,
			rating,
			bio,
			consultation_fee,
			phone,
			profile_image,
			available
		FROM experts
		WHERE id = $1;
	`

	err := s.db.GetContext(ctx, &expert, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &expert, nil
}

// GetExpertsByCategory returns experts by specialization.
func (s *expertStore) GetExpertsByCategory(ctx context.Context, category string) ([]models.Expert, error) {

	var experts []models.Expert

	query := `
		SELECT
			id,
			name,
			specialization,
			experience,
			language,
			rating,
			bio,
			consultation_fee,
			phone,
			profile_image,
			available
		FROM experts
		WHERE LOWER(specialization) = LOWER($1)
		ORDER BY rating DESC;
	`

	err := s.db.SelectContext(ctx, &experts, query, category)
	if err != nil {
		return nil, err
	}

	return experts, nil
}

// BookConsultation creates a consultation booking.
func (s *expertStore) BookConsultation(ctx context.Context, req *models.BookConsultationRequest) error {

	query := `
		INSERT INTO expert_consultations
		(
			user_id,
			expert_id,
			consultation_date,
			consultation_time,
			mode,
			problem,
			status
		)
		VALUES
		(
			$1,$2,$3,$4,$5,$6,'Pending'
		);
	`

	_, err := s.db.ExecContext(
		ctx,
		query,
		req.UserID,
		req.ExpertID,
		req.Date,
		req.Time,
		req.Mode,
		req.Problem,
	)

	return err
}

// GetConsultationHistory returns consultation history of a user.
func (s *expertStore) GetConsultationHistory(ctx context.Context, userID int64) ([]models.ConsultationHistory, error) {

	var history []models.ConsultationHistory

	query := `
		SELECT
			c.id AS booking_id,
			e.name AS expert_name,
			c.consultation_date AS date,
			c.consultation_time AS time,
			c.mode,
			c.status
		FROM expert_consultations c
		INNER JOIN experts e
			ON c.expert_id = e.id
		WHERE c.user_id = $1
		ORDER BY c.consultation_date DESC,
		         c.consultation_time DESC;
	`

	err := s.db.SelectContext(ctx, &history, query, userID)
	if err != nil {
		return nil, err
	}

	return history, nil
}
