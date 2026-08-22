package handler

import (
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/banner/controller"
	"kisaanSathi/pkg/services/banner/db"

	"github.com/gin-gonic/gin"
)

type handler struct {
	controller controller.BannerController
}

type BannerHandler interface {

	GetBanners(c *gin.Context)

	GetBannerByID(c *gin.Context)

	CreateBanner(c *gin.Context)

	UpdateBanner(c *gin.Context)

	DeleteBanner(c *gin.Context)

	GetActiveBanners(c *gin.Context)
}

func NewBannerHandler(controller controller.BannerController) BannerHandler {
	return &handler{
		controller: controller,
	}
}

func NewBannerController(repo repo.DataObject) controller.BannerController {
	store := db.NewDBObject(repo.Databases.PgDB)
	return controller.NewBannerController(store)
}