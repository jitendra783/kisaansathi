package controller

import (
	"kisaanSathi/pkg/services/mandi/models"
	"strconv"
)

// GetMandiBhav returns latest mandi prices
func (c *mandiController) GetMandiBhav() ([]models.MandiPriceResponse, error) {

	result, err := c.store.GetMandiBhav()
	if err != nil {
		return nil, err
	}

	response := make([]models.MandiPriceResponse, 0, len(result))

	for _, item := range result {
		// convert to int
		id, _ := strconv.Atoi(item.ID.String)
		response = append(response, models.MandiPriceResponse{
			ID:          id,
			Crop:        item.Crop.String,
			Market:      item.Market.String,
			State:       item.State.String,
			District:    item.District.String,
			MinPrice:    item.MinPrice.Float64,
			MaxPrice:    item.MaxPrice.Float64,
			ModalPrice:  item.ModalPrice.Float64,
			ArrivalDate: item.ArrivalDate.Time.Format("2006-01-02"),
		})
	}

	return response, nil
}

// GetMandiPrices returns all mandi prices
func (c *mandiController) GetMandiPrices() ([]models.MandiPriceResponse, error) {

	result, err := c.store.GetMandiPrices()
	if err != nil {
		return nil, err
	}

	response := make([]models.MandiPriceResponse, 0, len(result))

	for _, item := range result {
		id, _ := strconv.Atoi(item.ID.String)
		response = append(response, models.MandiPriceResponse{
			ID:          id,
			Crop:        item.Crop.String,
			Market:      item.Market.String,
			State:       item.State.String,
			District:    item.District.String,
			MinPrice:    item.MinPrice.Float64,
			MaxPrice:    item.MaxPrice.Float64,
			ModalPrice:  item.ModalPrice.Float64,
		})
	}

	return response, nil
}

// GetCropPrices returns prices for a crop
func (c *mandiController) GetCropPrices(crop string) ([]models.MandiPriceResponse, error) {

	result, err := c.store.GetCropPrices(crop)
	if err != nil {
		return nil, err
	}

	response := make([]models.MandiPriceResponse, 0, len(result))

	for _, item := range result {
		id, _ := strconv.Atoi(item.ID.String)
		response = append(response, models.MandiPriceResponse{
			ID:          id,
			Crop:        item.Crop.String,
			Market:      item.Market.String,
			State:       item.State.String,
			District:    item.District.String,
			MinPrice:    item.MinPrice.Float64,
			MaxPrice:    item.MaxPrice.Float64,
			ModalPrice:  item.ModalPrice.Float64,
		})
	}

	return response, nil
}

// GetStatePrices returns prices by state
func (c *mandiController) GetStatePrices(state string) ([]models.MandiPriceResponse, error) {

	result, err := c.store.GetStatePrices(state)
	if err != nil {
		return nil, err
	}

	response := make([]models.MandiPriceResponse, 0, len(result))

	for _, item := range result {
		id, _ := strconv.Atoi(item.ID.String)

		response = append(response, models.MandiPriceResponse{
			ID:          id,
			Crop:        item.Crop.String,
			Market:      item.Market.String,
			State:       item.State.String,
			District:    item.District.String,
			MinPrice:    item.MinPrice.Float64,
			MaxPrice:    item.MaxPrice.Float64,
			ModalPrice:  item.ModalPrice.Float64,
		})
	}

	return response, nil
}

// GetDistrictPrices returns prices by district
func (c *mandiController) GetDistrictPrices(district string) ([]models.MandiPriceResponse, error) {

	result, err := c.store.GetDistrictPrices(district)
	if err != nil {
		return nil, err
	}

	response := make([]models.MandiPriceResponse, 0, len(result))

	for _, item := range result {
		id, _ := strconv.Atoi(item.ID.String)
		response = append(response, models.MandiPriceResponse{
			ID:          id,
			Crop:        item.Crop.String,
			Market:      item.Market.String,
			State:       item.State.String,
			District:    item.District.String,
			MinPrice:    item.MinPrice.Float64,
			MaxPrice:    item.MaxPrice.Float64,
			ModalPrice:  item.ModalPrice.Float64,
		})
	}

	return response, nil
}

// GetTrendingPrices returns trending mandi prices
func (c *mandiController) GetTrendingPrices() ([]models.TrendingPriceResponse, error) {

	result, err := c.store.GetTrendingPrices()
	if err != nil {
		return nil, err
	}

	response := make([]models.TrendingPriceResponse, 0, len(result))

	for _, item := range result {
		response = append(response, models.TrendingPriceResponse{
			Crop:        item.Crop.String,
			Market:      item.Market.String,
			ModalPrice:  item.ModalPrice.Float64,
			PriceChange: item.Change.Float64,
			Trend:       item.Trend.String,
		})
	}

	return response, nil
}
func (c *mandiController) GetPriceComparison(crop, market string) (*models.PriceComparisonResponse, error) {

	result, err := c.store.GetPriceComparison(crop, market)
	if err != nil {
		return nil, err
	}

	resp := &models.PriceComparisonResponse{
		Crop:   crop,
		Market: market,
	}

	if len(result) > 0 {
		resp.CurrentPrice = result[0].ModalPrice
	}

	for _, item := range result {

		if len(resp.Last7Days) < 7 {
			resp.Last7Days = append(resp.Last7Days, item)
		}

		if len(resp.Last10Days) < 10 {
			resp.Last10Days = append(resp.Last10Days, item)
		}

		if len(resp.Last30Days) < 30 {
			resp.Last30Days = append(resp.Last30Days, item)
		}
	}

	return resp, nil
}
