package controller

import "kisaanSathi/pkg/services/mandi/models"

// GetMandiBhav returns latest mandi prices
func (c *mandiController) GetMandiBhav() ([]models.MandiPriceResponse, error) {

	result, err := c.store.GetMandiBhav()
	if err != nil {
		return nil, err
	}

	response := make([]models.MandiPriceResponse, 0, len(result))

	for _, item := range result {
		response = append(response, models.MandiPriceResponse{
			Crop:       item.Crop,
			Market:     item.Market,
			State:      item.State,
			District:   item.District,
			MinPrice:   item.MinPrice,
			MaxPrice:   item.MaxPrice,
			ModalPrice: item.ModalPrice,
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
		response = append(response, models.MandiPriceResponse{
			Crop:       item.Crop,
			Market:     item.Market,
			State:      item.State,
			District:   item.District,
			MinPrice:   item.MinPrice,
			MaxPrice:   item.MaxPrice,
			ModalPrice: item.ModalPrice,
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
		response = append(response, models.MandiPriceResponse{
			Crop:       item.Crop,
			Market:     item.Market,
			State:      item.State,
			District:   item.District,
			MinPrice:   item.MinPrice,
			MaxPrice:   item.MaxPrice,
			ModalPrice: item.ModalPrice,
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
		response = append(response, models.MandiPriceResponse{
			Crop:       item.Crop,
			Market:     item.Market,
			State:      item.State,
			District:   item.District,
			MinPrice:   item.MinPrice,
			MaxPrice:   item.MaxPrice,
			ModalPrice: item.ModalPrice,
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
		response = append(response, models.MandiPriceResponse{
			Crop:       item.Crop,
			Market:     item.Market,
			State:      item.State,
			District:   item.District,
			MinPrice:   item.MinPrice,
			MaxPrice:   item.MaxPrice,
			ModalPrice: item.ModalPrice,
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
			Crop:        item.Crop,
			Market:      item.Market,
			ModalPrice:  item.ModalPrice,
			PriceChange: item.Change,
			Trend:       item.Trend,
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