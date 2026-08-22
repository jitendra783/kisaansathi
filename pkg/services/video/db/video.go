package db

import "kisaanSathi/pkg/services/video/models"

func (s *videoStore) GetVideos() ([]models.Video, error) {

	var videos []models.Video

	err := s.db.Select(
		&videos,
		`SELECT
			id,
			title,
			description,
			thumbnail,
			video_url,
			category,
			duration,
			language,
			created_at
		FROM metadata_video
		ORDER BY created_at DESC`,
	)

	if err != nil {
		return nil, err
	}

	return videos, nil
}
func (s *videoStore) GetVideo(id int64) (*models.Video, error) {

	var video models.Video

	err := s.db.Get(
		&video,
		`SELECT
			id,
			title,
			description,
			thumbnail,
			video_url,
			category,
			duration,
			language,
			created_at
		FROM metadata_video
		WHERE id=$1`,
		id,
	)

	if err != nil {
		return nil, err
	}

	return &video, nil
}
func (s *videoStore) GetVideosByCategory(category string) ([]models.Video, error) {

	var videos []models.Video

	err := s.db.Select(
		&videos,
		`SELECT
			id,
			title,
			description,
			thumbnail,
			video_url,
			category,
			duration,
			language,
			created_at
		FROM metadata_video
		WHERE category=$1
		ORDER BY created_at DESC`,
		category,
	)

	if err != nil {
		return nil, err
	}

	return videos, nil
}