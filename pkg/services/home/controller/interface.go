package controller

import (
	"context"
	"kisaanSathi/pkg/services/home/db"
	"kisaanSathi/pkg/services/home/models"
)

type homeController struct {
	store db.HomeStore
}

type HomeController interface {
	GetHome(context.Context) (*models.HomeResponse, error)
	GetDashboard(context.Context) (*models.DashboardResponse, error)
}

func NewHomeController(store db.HomeStore) HomeController {
	return &homeController{
		store: store,
	}
}
