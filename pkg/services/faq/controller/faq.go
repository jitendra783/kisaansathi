package controller

import (
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/faq/models"
)

func (c *faqController) GetFAQs() (*models.FAQResponse, error) {

	logger.Log().Info("FAQ Controller: GetFAQs started")

	result, err := c.store.GetFAQs()
	if err != nil {
		logger.Log().Error("FAQ Controller: " + err.Error())
		return nil, err
	}

	logger.Log().Sugar().Infof("FAQ Controller: Total FAQs fetched: %d", len(result))

	return &models.FAQResponse{
		Data: result,
	}, nil
}

func (c *faqController) GetFAQByID(id int64) (*models.FAQDetailResponse, error) {

	logger.Log().Sugar().Infof("FAQ Controller: GetFAQByID called with ID=%d", id)

	result, err := c.store.GetFAQByID(id)
	if err != nil {
		logger.Log().Error("FAQ Controller: " + err.Error())
		return nil, err
	}

	logger.Log().Info("FAQ Controller: FAQ fetched successfully")

	return &models.FAQDetailResponse{
		Data: *result,
	}, nil
}

func (c *faqController) GetFAQsByCategory(category string) (*models.FAQResponse, error) {

	logger.Log().Info("FAQ Controller: GetFAQsByCategory started")
	logger.Log().Sugar().Infof("Category: %s", category)

	result, err := c.store.GetFAQsByCategory(category)
	if err != nil {
		logger.Log().Error("FAQ Controller: " + err.Error())
		return nil, err
	}

	logger.Log().Sugar().Infof("Total FAQs fetched for category '%s': %d", category, len(result))

	return &models.FAQResponse{
		Data: result,
	}, nil
}