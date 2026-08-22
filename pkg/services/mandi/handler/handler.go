package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *mandiHandler) GetMandiBhav(c *gin.Context) {

	response, err := h.controller.GetMandiBhav()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
func (h *mandiHandler) GetMandiPrices(c *gin.Context) {

	response, err := h.controller.GetMandiPrices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
func (h *mandiHandler) GetCropPrices(c *gin.Context) {

	crop := c.Param("crop")

	response, err := h.controller.GetCropPrices(crop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
func (h *mandiHandler) GetStatePrices(c *gin.Context) {

	state := c.Param("state")

	response, err := h.controller.GetStatePrices(state)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
func (h *mandiHandler) GetDistrictPrices(c *gin.Context) {

	district := c.Param("district")

	response, err := h.controller.GetDistrictPrices(district)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
func (h *mandiHandler) GetTrendingPrices(c *gin.Context) {

	response, err := h.controller.GetTrendingPrices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
func (h *mandiHandler) GetPriceComparison(c *gin.Context) {

	crop := c.Param("crop")
	market := c.Param("market")

	response, err := h.controller.GetPriceComparison(crop, market)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}