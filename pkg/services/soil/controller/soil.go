package controller

import (
	"errors"
	"kisaanSathi/pkg/services/soil/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *soilController) GetSoilTypes(ctx *gin.Context) (*models.SoilTypesResponse, error) {

	result, err := c.store.GetSoilTypes()
	if err != nil {
		return nil, err
	}

	return &models.SoilTypesResponse{
		Data: result,
	}, nil
}
func (c *soilController) GetSoilReport(ctx *gin.Context) (*models.SoilReportResponse, error) {

	userIDStr := ctx.Query("user_id")
	if userIDStr == "" {
		return nil, errors.New("user_id is required")
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return nil, errors.New("invalid user_id")
	}

	report, err := c.store.GetSoilReport(userID)
	if err != nil {
		return nil, err
	}

	if report == nil {
		return &models.SoilReportResponse{}, nil
	}

	return &models.SoilReportResponse{
		Data: report,
	}, nil
}

func (c *soilController) CreateSoilTest(ctx *gin.Context) (*models.CommonResponse, error) {

	var req models.CreateSoilTestRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	err := c.store.CreateSoilTest(&req)
	if err != nil {
		return nil, err
	}

	return &models.CommonResponse{
		Message: "Soil test created successfully",
	}, nil
}

func (c *soilController) GetSoilRecommendation(ctx *gin.Context) (*models.SoilRecommendationResponse, error) {

	userIDStr := ctx.Query("user_id")
	if userIDStr == "" {
		return nil, errors.New("user_id is required")
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return nil, errors.New("invalid user_id")
	}

	result, err := c.store.GetSoilRecommendation(userID)
	if err != nil {
		return nil, err
	}

	return &models.SoilRecommendationResponse{
		Data: result,
	}, nil
}
