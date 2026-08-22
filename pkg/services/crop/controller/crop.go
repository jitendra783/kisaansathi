package controller

import "kisaanSathi/pkg/services/crop/models"

func (c *cropController) GetCrops() (*models.CropListResponse, error) {

	result, err := c.store.GetCrops()
	if err != nil {
		return nil, err
	}

	return &models.CropListResponse{
		Data: result,
	}, nil
}
func (c *cropController) GetCrop(id int64) (*models.CropResponse, error) {

	result, err := c.store.GetCrop(id)
	if err != nil {
		return nil, err
	}

	return &models.CropResponse{
		Data: result,
	}, nil
}
func (c *cropController) CreateCrop(req models.CreateCropRequest) error {

	return c.store.CreateCrop(req)
}
func (c *cropController) UpdateCrop(id int64, req models.UpdateCropRequest) error {

	return c.store.UpdateCrop(id, req)
}
func (c *cropController) DeleteCrop(id int64) error {

	return c.store.DeleteCrop(id)
}
func (c *cropController) GetCropSeason(season string) (*models.CropListResponse, error) {

	result, err := c.store.GetCropSeason(season)
	if err != nil {
		return nil, err
	}

	return &models.CropListResponse{
		Data: result,
	}, nil
}

func (c *cropController) GetRecommendedCrops(soil, district string) (*models.RecommendedCropResponse, error) {

	result, err := c.store.GetRecommendedCrops(soil, district)
	if err != nil {
		return nil, err
	}

	return &models.RecommendedCropResponse{
		Data: result,
	}, nil
}