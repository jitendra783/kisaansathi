package db

import "kisaanSathi/pkg/services/crop/models"

func (s *cropStore) GetCrops() ([]models.Crop, error) {

	var crops []models.Crop

	err := s.db.Select(&crops, `
		SELECT
			id,
			name,
			season,
			soil_type,
			duration,
			image_url,
			description,
			created_at
		FROM crop
		WHERE is_active = true
		ORDER BY name`)

	if err != nil {
		return nil, err
	}

	return crops, nil
}

func (s *cropStore) GetCrop(id int64) (*models.Crop, error) {

	var crop models.Crop

	err := s.db.Get(&crop, `
		SELECT
			id,
			name,
			season,
			soil_type,
			duration,
			image_url,
			description,
			created_at
		FROM crop
		WHERE id=$1
		  AND is_active=true`,
		id,
	)

	if err != nil {
		return nil, err
	}

	return &crop, nil
}

func (s *cropStore) CreateCrop(req models.CreateCropRequest) error {

	_, err := s.db.Exec(`
		INSERT INTO crop
		(
			name,
			season,
			soil_type,
			duration,
			image_url,
			description,
			is_active
		)
		VALUES
		(
			$1,$2,$3,$4,$5,$6,true
		)`,
		req.Name,
		req.Season,
		req.SoilType,
		req.Duration,
		req.ImageURL,
		req.Description,
	)

	return err
}

func (s *cropStore) UpdateCrop(id int64, req models.UpdateCropRequest) error {

	_, err := s.db.Exec(`
		UPDATE crop
		SET
			name=$1,
			season=$2,
			soil_type=$3,
			duration=$4,
			image_url=$5,
			description=$6,
			updated_at=NOW()
		WHERE id=$7`,
		req.Name,
		req.Season,
		req.SoilType,
		req.Duration,
		req.ImageURL,
		req.Description,
		id,
	)

	return err
}

func (s *cropStore) DeleteCrop(id int64) error {

	_, err := s.db.Exec(`
		UPDATE crop
		SET
			is_active=false,
			updated_at=NOW()
		WHERE id=$1`,
		id,
	)

	return err
}

func (s *cropStore) GetCropSeason(season string) ([]models.Crop, error) {

	var crops []models.Crop

	err := s.db.Select(&crops, `
		SELECT
			id,
			name,
			season,
			soil_type,
			duration,
			image_url,
			description,
			created_at
		FROM crop
		WHERE season=$1
		  AND is_active=true
		ORDER BY name`,
		season,
	)

	if err != nil {
		return nil, err
	}

	return crops, nil
}

func (s *cropStore) GetRecommendedCrops(soil, district string) ([]models.RecommendedCrop, error) {

	var crops []models.RecommendedCrop

	err := s.db.Select(&crops, `
		SELECT
			name AS crop,
			confidence
		FROM crop_recommendation
		WHERE soil_type=$1
		  AND district=$2
		ORDER BY confidence DESC`,
		soil,
		district,
	)

	if err != nil {
		return nil, err
	}

	return crops, nil
}
