package handler

import (
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/soil/controller"
	"kisaanSathi/pkg/services/soil/db"

	"github.com/gin-gonic/gin"
)

type soilHandler struct {
	controller controller.SoilController
}
type SoilHandler interface {
	GetSoilTypes(ctx *gin.Context)
	GetSoilReport(ctx *gin.Context)
	CreateSoilTest(ctx *gin.Context)
	GetSoilRecommendation(ctx *gin.Context)
}

func NewSoilHandler(controller controller.SoilController) SoilHandler {
	return &soilHandler{
		controller: controller,
	}
}
func SoilController(repo repo.DataObject) controller.SoilController {
	store := db.NewSoilStore(repo.Databases.PgDB)
	return controller.NewSoilController(store)
}
