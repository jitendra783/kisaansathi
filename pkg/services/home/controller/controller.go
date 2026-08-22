package controller

import (
	"context"
	"kisaanSathi/pkg/services/home/models"
)

func (c *homeController) GetHome(ctx context.Context) (*models.HomeResponse, error) {
	return c.store.GetHomeData(ctx)
}
func (c *homeController) GetDashboard(ctx context.Context) (*models.DashboardResponse, error) {
	return c.store.GetDashboard(ctx)
}
