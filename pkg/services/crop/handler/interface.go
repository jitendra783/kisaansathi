package handler

import (
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/crop/controller"
	"kisaanSathi/pkg/services/crop/db"

	"github.com/gin-gonic/gin"
)

type cropHandler struct {
	controller controller.CropController
}

type CropHandler interface {
	GetCrops(ctx *gin.Context)
	GetCrop(ctx *gin.Context)
	CreateCrop(ctx *gin.Context)
	UpdateCrop(ctx *gin.Context)
	DeleteCrop(ctx *gin.Context)

	GetCropSeason(ctx *gin.Context)
	GetRecommendedCrops(ctx *gin.Context)
}

func NewCropHandler(controller controller.CropController) CropHandler {
	return &cropHandler{
		controller: controller,
	}
}
func NewCropController(repo repo.DataObject) controller.CropController {
	store := db.NewCropStore(repo.Databases.PgDB)
	return controller.NewCropController(store)
}
