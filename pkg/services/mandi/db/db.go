package db

import "kisaanSathi/pkg/services/mandi/models"

func (d *mandiStore) GetMandiBhav() ([]models.MandiPrice, error) {

	var prices []models.MandiPrice

	query := `
		SELECT
			id,
			crop,
			variety,
			market,
			district,
			state,
			min_price,
			max_price,
			modal_price,
			arrival_date
		FROM market_prices
		ORDER BY arrival_date DESC
	`

	err := d.db.Select(&prices, query)
	if err != nil {
		return nil, err
	}

	return prices, nil
}

func (d *mandiStore) GetMandiPrices() ([]models.MandiPrice, error) {

	var prices []models.MandiPrice

	query := `
		SELECT
			id,
			crop,
			variety,
			market,
			district,
			state,
			min_price,
			max_price,
			modal_price,
			arrival_date
		FROM market_prices
		ORDER BY crop, market
	`

	err := d.db.Select(&prices, query)
	if err != nil {
		return nil, err
	}

	return prices, nil
}

func (d *mandiStore) GetCropPrices(crop string) ([]models.MandiPrice, error) {

	var prices []models.MandiPrice

	query := `
		SELECT
			id,
			crop,
			variety,
			market,
			district,
			state,
			min_price,
			max_price,
			modal_price,
			arrival_date
		FROM market_prices
		WHERE LOWER(crop) = LOWER($1)
		ORDER BY modal_price DESC
	`

	err := d.db.Select(&prices, query, crop)
	if err != nil {
		return nil, err
	}

	return prices, nil
}
func (d *mandiStore) GetStatePrices(state string) ([]models.MandiPrice, error) {

	var prices []models.MandiPrice

	query := `
		SELECT
			id,
			crop,
			variety,
			market,
			district,
			state,
			min_price,
			max_price,
			modal_price,
			arrival_date
		FROM market_prices
		WHERE LOWER(state) = LOWER($1)
		ORDER BY market
	`

	err := d.db.Select(&prices, query, state)
	if err != nil {
		return nil, err
	}

	return prices, nil
}

func (d *mandiStore) GetDistrictPrices(district string) ([]models.MandiPrice, error) {

	var prices []models.MandiPrice

	query := `
		SELECT
			id,
			crop,
			variety,
			market,
			district,
			state,
			min_price,
			max_price,
			modal_price,
			arrival_date
		FROM market_prices
		WHERE LOWER(district) = LOWER($1)
		ORDER BY market
	`

	err := d.db.Select(&prices, query, district)
	if err != nil {
		return nil, err
	}

	return prices, nil
}

func (d *mandiStore) GetTrendingPrices() ([]models.TrendingPrice, error) {

	var prices []models.TrendingPrice

	query := `
		SELECT
			crop,
			market,
			modal_price,
			change,
			trend
		FROM market_prices
		ORDER BY change DESC
		LIMIT 10
	`

	err := d.db.Select(&prices, query)
	if err != nil {
		return nil, err
	}

	return prices, nil
}
func (d *mandiStore) GetPriceComparison(crop, market string) ([]models.PriceHistory, error) {

	var prices []models.PriceHistory

	query := `
		SELECT
			arrival_date,
			min_price,
			max_price,
			modal_price
		FROM market_prices
		WHERE LOWER(crop) = LOWER($1)
		  AND LOWER(market) = LOWER($2)
		  AND arrival_date >= CURRENT_DATE - INTERVAL '30 days'
		ORDER BY arrival_date DESC
	`

	err := d.db.Select(&prices, query, crop, market)
	if err != nil {
		return nil, err
	}

	return prices, nil
}
