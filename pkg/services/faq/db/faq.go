package db

import (
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/faq/models"
)

func (s *faqStore) GetFAQs() ([]models.FAQ, error) {

	logger.Log().Info("FAQ DB: GetFAQs started")

	faqs := make([]models.FAQ, 0)

	err := s.db.Select(
		&faqs,
		`SELECT
			id,
			question,
			answer
		FROM ui_faq
		ORDER BY id`,
	)

	if err != nil {
		logger.Log().Error("FAQ DB: Failed to fetch FAQs: " + err.Error())
		return nil, err
	}

	logger.Log().Sugar().Infof("FAQ DB: Total FAQs fetched: %d", len(faqs))

	return faqs, nil
}

func (s *faqStore) GetFAQByID(id int64) (*models.FAQ, error) {

	logger.Log().Sugar().Infof("FAQ DB: GetFAQByID called with ID=%d", id)

	var faq models.FAQ

	err := s.db.Get(
		&faq,
		`SELECT
			id,
			question,
			answer
		FROM ui_faq
		WHERE id=$1`,
		id,
	)

	if err != nil {
		logger.Log().Error("FAQ DB: Failed to fetch FAQ: " + err.Error())
		return nil, err
	}

	logger.Log().Info("FAQ DB: FAQ fetched successfully")

	return &faq, nil
}

func (s *faqStore) GetFAQsByCategory(category string) ([]models.FAQ, error) {

	logger.Log().Sugar().Infof("FAQ DB: GetFAQsByCategory called. Category=%s", category)

	faqs := make([]models.FAQ, 0)

	err := s.db.Select(
		&faqs,
		`SELECT
			id,
			question,
			answer
		FROM ui_faq
		WHERE category=$1
		ORDER BY id`,
		category,
	)

	if err != nil {
		logger.Log().Error("FAQ DB: Failed to fetch category FAQs: " + err.Error())
		return nil, err
	}

	logger.Log().Sugar().Infof("FAQ DB: Total FAQs for category '%s': %d", category, len(faqs))

	return faqs, nil
}