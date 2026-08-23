package db

import (
	"database/sql"
	"kisaanSathi/pkg/services/soil/models"
)

func (s *soilStore) GetSoilTypes() ([]models.SoilType, error) {

	var result []models.SoilType

	query := `
		SELECT
			id,
			name,
			description
		FROM metadata_soil_types
		ORDER BY name;
	`

	err := s.db.Select(&result, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (s *soilStore) GetSoilReport(userID int64) (*models.SoilReport, error) {

	var report models.SoilReport

	query := `
		SELECT
			id,
			user_id,
			soil_type,
			ph,
			nitrogen,
			phosphorus,
			potassium,
			organic_carbon,
			electrical_ec,
			moisture,
			status,
			created_at
		FROM soil_reports
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT 1;
	`

	err := s.db.Get(&report, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &report, nil
}

func (s *soilStore) CreateSoilTest(req *models.CreateSoilTestRequest) error {

	query := `
		INSERT INTO soil_reports
		(
			user_id,
			soil_type,
			ph,
			nitrogen,
			phosphorus,
			potassium,
			organic_carbon,
			electrical_ec,
			moisture,
			status,
			created_at
		)
		VALUES
		(
			:user_id,
			:soil_type,
			:ph,
			:nitrogen,
			:phosphorus,
			:potassium,
			:organic_carbon,
			:electrical_ec,
			:moisture,
			'Pending',
			NOW()
		)
	`

	_, err := s.db.NamedExec(query, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *soilStore) GetSoilRecommendation(userID int64) ([]models.SoilRecommendation, error) {

	var recommendations []models.SoilRecommendation

	query := `
		SELECT
			r.id,
			r.crop,
			r.soil_type,
			r.recommendation,
			r.fertilizer,
			r.irrigation
		FROM soil_recommendations r
		INNER JOIN soil_reports sr
			ON sr.soil_type = r.soil_type
		WHERE sr.user_id = $1
		ORDER BY r.crop;
	`

	err := s.db.Select(&recommendations, query, userID)
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}
