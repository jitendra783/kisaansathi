package controller

import "kisaanSathi/pkg/services/video/models"

func (c *videoController) GetVideos() (*models.VideosResponse, error) {

	result, err := c.store.GetVideos()
	if err != nil {
		return nil, err
	}

	return &models.VideosResponse{
		Data: result,
	}, nil
}

func (c *videoController) GetVideo(id int64) (*models.VideoResponse, error) {

	result, err := c.store.GetVideo(id)
	if err != nil {
		return nil, err
	}

	return &models.VideoResponse{
		Data: result,
	}, nil
}
func (c *videoController) GetVideosByCategory(category string) (*models.VideosResponse, error) {

	result, err := c.store.GetVideosByCategory(category)
	if err != nil {
		return nil, err
	}

	return &models.VideosResponse{
		Data: result,
	}, nil
}
