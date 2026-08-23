package db

import (
	"context"
	"kisaanSathi/pkg/services/home/models"
)

func (s *homeStore) GetHomeData(ctx context.Context) (*models.HomeResponse, error) {

	var banners []models.Banner
	var cards []models.HomeCard

	err := s.db.Select(&banners,
		`SELECT id,title,image_url,redirect_url,order_no
		 FROM metadata_banner
		 ORDER BY order_no`)

	if err != nil {
		return nil, err
	}

	err = s.db.Select(&cards,
		`SELECT id,title,icon,redirect_url,order_no
		 FROM metadata_home_cards
		 ORDER BY order_no`)

	if err != nil {
		return nil, err
	}

	return &models.HomeResponse{
		Banners: banners,
		Cards:   cards,
	}, nil
}

func (s *homeStore) GetDashboard(ctx context.Context) (*models.DashboardResponse, error) {

	resp := &models.DashboardResponse{
		Weather: nil,
		Mandi:   nil,
		Feeds:   nil,
		User:    nil,
	}

	return resp, nil
}
