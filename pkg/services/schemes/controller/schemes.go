package controller

import (
	"errors"
	"math"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/schemes/models"
)

// GetSchemes returns paginated schemes.
func (c *schemeController) GetSchemes(page, limit int) (*models.SchemeResponse, error) {

	logger.Log().Sugar().Infof(
		"Scheme Controller: GetSchemes called [page=%d limit=%d]",
		page,
		limit,
	)

	result, total, err := c.store.GetSchemes(page, limit)
	if err != nil {
		logger.Log().Error(err.Error())
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	logger.Log().Sugar().Infof(
		"Scheme Controller: Total=%d Returned=%d TotalPages=%d",
		total,
		len(result),
		totalPages,
	)

	for i, scheme := range result {
		logger.Log().Sugar().Infof(
			"Scheme %d: %+v",
			i+1,
			scheme,
		)
	}

	return &models.SchemeResponse{
		Schemes: result,
		Pagination: models.Pagination{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

// GetScheme returns scheme details by ID.
func (c *schemeController) GetScheme(id int64) (*models.SchemeDetailResponse, error) {

	logger.Log().Sugar().Infof(
		"Scheme Controller: GetScheme called [id=%d]",
		id,
	)

	result, err := c.store.GetScheme(id)
	if err != nil {
		logger.Log().Error(err.Error())
		return nil, err
	}

	if result == nil {
		logger.Log().Warn("Scheme Controller: Scheme not found")
		return nil, errors.New("scheme not found")
	}

	logger.Log().Sugar().Infof(
		"Scheme Controller: Scheme fetched: %+v",
		*result,
	)

	return &models.SchemeDetailResponse{
		Scheme: result,
	}, nil
}

// GetEligibleSchemes returns eligible schemes.
func (c *schemeController) GetEligibleSchemes(
	state,
	category,
	farmerType string,
) (*models.SchemeResponse, error) {

	logger.Log().Sugar().Infof(
		"Scheme Controller: GetEligibleSchemes [state=%s category=%s farmerType=%s]",
		state,
		category,
		farmerType,
	)

	result, err := c.store.GetEligibleSchemes(
		state,
		category,
		farmerType,
	)
	if err != nil {
		logger.Log().Error(err.Error())
		return nil, err
	}

	logger.Log().Sugar().Infof(
		"Eligible schemes fetched: %d",
		len(result),
	)

	for i, scheme := range result {
		logger.Log().Sugar().Infof(
			"Eligible Scheme %d: %+v",
			i+1,
			scheme,
		)
	}

	return &models.SchemeResponse{
		Schemes: result,
	}, nil
}

// GetStateSchemes returns schemes for a particular state.
func (c *schemeController) GetStateSchemes(
	state string,
) (*models.SchemeResponse, error) {

	logger.Log().Sugar().Infof(
		"Scheme Controller: GetStateSchemes [state=%s]",
		state,
	)

	result, err := c.store.GetStateSchemes(state)
	if err != nil {
		logger.Log().Error(err.Error())
		return nil, err
	}

	logger.Log().Sugar().Infof(
		"State schemes fetched: %d",
		len(result),
	)

	for i, scheme := range result {
		logger.Log().Sugar().Infof(
			"State Scheme %d: %+v",
			i+1,
			scheme,
		)
	}

	return &models.SchemeResponse{
		Schemes: result,
	}, nil
}
