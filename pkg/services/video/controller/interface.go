package controller

import (
	"kisaanSathi/pkg/services/video/db"
	"kisaanSathi/pkg/services/video/models"
)

type VideoController interface {
	GetVideos() (*models.VideosResponse, error)
	GetVideo(id int64) (*models.VideoResponse, error)
	GetVideosByCategory(category string) (*models.VideosResponse, error)
}
type videoController struct {
	store db.VideoStore
}

func NewVideoController(store db.VideoStore) VideoController {
	return &videoController{
		store: store,
	}
}
