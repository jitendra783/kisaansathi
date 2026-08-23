package handler

import (
	"kisaanSathi/pkg/repo"

	"kisaanSathi/pkg/services/mandi/controller"
	"kisaanSathi/pkg/services/mandi/db"

	"github.com/gin-gonic/gin"
)

type mandiHandler struct {
	controller controller.MandiController
}
type MandiHandler interface {
	GetMandiBhav(*gin.Context)
	GetMandiPrices(*gin.Context)
	GetCropPrices(*gin.Context)
	GetStatePrices(*gin.Context)
	GetDistrictPrices(*gin.Context)
	GetTrendingPrices(*gin.Context)
	GetPriceComparison(*gin.Context)
}

func NewMandiHandler(controller controller.MandiController) MandiHandler {
	return &mandiHandler{
		controller: controller,
	}
}
func NewMandiController(repo repo.DataObject) controller.MandiController {
	store := db.NewMandiStore(repo.Databases.PgDB)
	return controller.NewMandiController(store)
}
