package db

import (
	"kisaanSathi/pkg/services/video/models"

	"github.com/jmoiron/sqlx"
)

type videoStore struct {
	db *sqlx.DB
}
type VideoStore interface {
	GetVideos() ([]models.Video, error)
	GetVideo(id int64) (*models.Video, error)
	GetVideosByCategory(category string) ([]models.Video, error)
}

func NewVideoStore(db *sqlx.DB) VideoStore {
	return &videoStore{
		db: db,
	}
}
