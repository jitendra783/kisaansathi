package db

import (
	"database/sql"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/schemes/models"
)

// GetSchemes returns paginated active schemes.
func (s *schemeStore) GetSchemes(page, limit int) ([]models.Scheme, int, error) {

	logger.Log().Sugar().Infof(
		"Scheme DB: GetSchemes called [page=%d limit=%d]",
		page,
		limit,
	)

	offset := (page - 1) * limit

	var (
		schemes []models.Scheme
		total   int
	)

	countQuery := `
		SELECT COUNT(*)
		FROM schemes
		WHERE status = 'ACTIVE';
	`

	if err := s.db.Get(&total, countQuery); err != nil {
		logger.Log().Error(err.Error())
		return nil, 0, err
	}

	query := `
		SELECT
			id,
			name,
			description,
			category,
			state,
			eligibility,
			benefits,
			official_url,
			application_mode,
			status
		FROM schemes
		WHERE status = 'ACTIVE'
		ORDER BY name
		LIMIT $1 OFFSET $2;
	`

	err := s.db.Select(
		&schemes,
		query,
		limit,
		offset,
	)
	if err != nil {
		logger.Log().Error(err.Error())
		return nil, 0, err
	}

	logger.Log().Sugar().Infof(
		"Scheme DB: Total=%d Returned=%d",
		total,
		len(schemes),
	)

	for i, scheme := range schemes {
		logger.Log().Sugar().Infof(
			"Scheme %d: %+v",
			offset+i+1,
			scheme,
		)
	}

	return schemes, total, nil
}

// GetScheme returns scheme details by ID.
func (s *schemeStore) GetScheme(id int64) (*models.Scheme, error) {

	logger.Log().Sugar().Infof(
		"Scheme DB: GetScheme called [id=%d]",
		id,
	)

	var scheme models.Scheme

	query := `
		SELECT
			id,
			name,
			description,
			category,
			state,
			eligibility,
			benefits,
			official_url,
			application_mode,
			status
		FROM schemes
		WHERE id = $1;
	`

	err := s.db.Get(&scheme, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Log().Warn("Scheme DB: No scheme found")
			return nil, nil
		}

		logger.Log().Error(err.Error())
		return nil, err
	}

	logger.Log().Sugar().Infof(
		"Scheme DB: Scheme fetched: %+v",
		scheme,
	)

	return &scheme, nil
}

// GetEligibleSchemes returns eligible schemes based on filters.
func (s *schemeStore) GetEligibleSchemes(
	state,
	category,
	farmerType string,
) ([]models.Scheme, error) {

	logger.Log().Sugar().Infof(
		"Scheme DB: GetEligibleSchemes [state=%s category=%s farmerType=%s]",
		state,
		category,
		farmerType,
	)

	var schemes []models.Scheme

	query := `
		SELECT
			id,
			name,
			description,
			category,
			state,
			eligibility,
			benefits,
			official_url,
			application_mode,
			status
		FROM schemes
		WHERE status = 'ACTIVE'
			AND ($1 = '' OR state = $1)
			AND ($2 = '' OR category = $2)
			AND ($3 = '' OR eligibility ILIKE '%' || $3 || '%')
		ORDER BY name;
	`

	err := s.db.Select(
		&schemes,
		query,
		state,
		category,
		farmerType,
	)
	if err != nil {
		logger.Log().Error(err.Error())
		return nil, err
	}

	logger.Log().Sugar().Infof(
		"Eligible Schemes Returned: %d",
		len(schemes),
	)

	return schemes, nil
}

// GetStateSchemes returns schemes for a particular state.
func (s *schemeStore) GetStateSchemes(state string) ([]models.Scheme, error) {

	logger.Log().Sugar().Infof(
		"Scheme DB: GetStateSchemes [state=%s]",
		state,
	)

	var schemes []models.Scheme

	query := `
		SELECT
			id,
			name,
			description,
			category,
			state,
			eligibility,
			benefits,
			official_url,
			application_mode,
			status
		FROM schemes
		WHERE status = 'ACTIVE'
			AND state = $1
		ORDER BY name;
	`

	err := s.db.Select(&schemes, query, state)
	if err != nil {
		logger.Log().Error(err.Error())
		return nil, err
	}

	logger.Log().Sugar().Infof(
		"State Schemes Returned: %d",
		len(schemes),
	)

	return schemes, nil
}