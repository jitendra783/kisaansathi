package controller

import (
	"context"
	"log"

	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/banner/models"
)

func (c *controller) GetBanners(ctx context.Context) ([]*models.BannerResponse, error) {

	logger.Log().Info("Banner Controller: GetBanners started")

	banners, err := c.store.GetBanners(ctx)
	if err != nil {
		logger.Log().Error(err.Error())
		return nil, err
	}

	response := make([]*models.BannerResponse, 0, len(banners))

	for _, banner := range banners {
		response = append(response, &models.BannerResponse{
			ID:           banner.ID,
			Page:         banner.Page,
			Title:        banner.Title,
			Subtitle:     banner.Subtitle,
			ImageURL:     banner.ImageURL,
			ActionURL:    banner.ActionURL,
			DisplayOrder: banner.DisplayOrder,
			IsActive:     banner.IsActive,
		})
	}

	logger.Log().Sugar().Infof("Banner Controller: Total banners=%d", len(response))

	return response, nil
}

func (c *controller) GetBannerByID(ctx context.Context, id int64) (*models.BannerResponse, error) {

	logger.Log().Sugar().Infof("Banner Controller: GetBannerByID ID=%d", id)

	banner, err := c.store.GetBannerByID(ctx, id)
	if err != nil {
		logger.Log().Error(err.Error())
		return nil, err
	}

	return &models.BannerResponse{
		ID:           banner.ID,
		Page:         banner.Page,
		Title:        banner.Title,
		Subtitle:     banner.Subtitle,
		ImageURL:     banner.ImageURL,
		ActionURL:    banner.ActionURL,
		DisplayOrder: banner.DisplayOrder,
		IsActive:     banner.IsActive,
	}, nil
}

func (c *controller) GetActiveBanners(ctx context.Context) ([]*models.BannerResponse, error) {

	logger.Log().Info("Banner Controller: GetActiveBanners started")

	banners, err := c.store.GetActiveBanners(ctx)
	if err != nil {
		logger.Log().Error(err.Error())
		return nil, err
	}

	response := make([]*models.BannerResponse, 0, len(banners))

	for _, banner := range banners {
		response = append(response, &models.BannerResponse{
			ID:           banner.ID,
			Page:         banner.Page,
			Title:        banner.Title,
			Subtitle:     banner.Subtitle,
			ImageURL:     banner.ImageURL,
			ActionURL:    banner.ActionURL,
			DisplayOrder: banner.DisplayOrder,
			IsActive:     banner.IsActive,
		})
	}
	log.Println("controller:", response)

	logger.Log().Sugar().Infof("Banner Controller: Active banners=%d", len(response))
	log.Println("controller:", response)
	return response, nil
}

func (c *controller) CreateBanner(ctx context.Context, req *models.CreateBannerRequest) error {

	logger.Log().Info("Banner Controller: CreateBanner")

	banner := &models.Banner{
		Page:         req.Page,
		Title:        req.Title,
		Subtitle:     req.Subtitle,
		ImageURL:     req.ImageURL,
		ActionURL:    req.ActionURL,
		DisplayOrder: req.DisplayOrder,
		IsActive:     req.IsActive,
	}

	return c.store.CreateBanner(ctx, banner)
}

func (c *controller) UpdateBanner(ctx context.Context, req *models.UpdateBannerRequest) error {

	logger.Log().Sugar().Infof("Banner Controller: UpdateBanner ID=%d", req.ID)

	banner := &models.Banner{
		ID:           req.ID,
		Page:         req.Page,
		Title:        req.Title,
		Subtitle:     req.Subtitle,
		ImageURL:     req.ImageURL,
		ActionURL:    req.ActionURL,
		DisplayOrder: req.DisplayOrder,
		IsActive:     req.IsActive,
	}

	return c.store.UpdateBanner(ctx, banner)
}

func (c *controller) DeleteBanner(ctx context.Context, id int64) error {

	logger.Log().Sugar().Infof("Banner Controller: DeleteBanner ID=%d", id)

	return c.store.DeleteBanner(ctx, id)
}
