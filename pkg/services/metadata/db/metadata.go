package db

import (
	"kisaanSathi/pkg/services/metadata/models"
)

func (s *metadataStore) GetMetadata() ([]models.MetadataDB, error) {

	query := `
		SELECT
			id,
			key,
			value,
			type,
			is_active,
			created_at,
			updated_at
		FROM metadata
		WHERE is_active = true
		ORDER BY id;
	`

	var metadata []models.MetadataDB

	err := s.db.Select(&metadata, query)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}
