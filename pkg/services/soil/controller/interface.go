package controller

import (
	"kisaanSathi/pkg/services/soil/db"
	"kisaanSathi/pkg/services/soil/models"

	"github.com/gin-gonic/gin"
)

type soilController struct {
	store db.SoilStore
}

// AddSoilDetails implements [SoilController].

type SoilController interface {
	GetSoilTypes(ctx *gin.Context) (*models.SoilTypesResponse, error)
	GetSoilReport(ctx *gin.Context) (*models.SoilReportResponse, error)
	CreateSoilTest(ctx *gin.Context) (*models.CommonResponse, error)
	GetSoilRecommendation(ctx *gin.Context) (*models.SoilRecommendationResponse, error)
}

func NewSoilController(store db.SoilStore) SoilController {
	return &soilController{
		store: store,
	}
}
