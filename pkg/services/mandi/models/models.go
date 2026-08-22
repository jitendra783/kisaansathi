package models

import "time"

type MandiPrice struct {
	ID          int64     `db:"id"`
	Crop        string    `db:"crop"`
	Variety     string    `db:"variety"`
	Market      string    `db:"market"`
	District    string    `db:"district"`
	State       string    `db:"state"`
	MinPrice    float64   `db:"min_price"`
	MaxPrice    float64   `db:"max_price"`
	ModalPrice  float64   `db:"modal_price"`
	ArrivalDate time.Time `db:"arrival_date"`
}

type MandiPriceResponse struct {
	ID          int64   `json:"id"`
	Crop        string  `json:"crop"`
	Variety     string  `json:"variety"`
	Market      string  `json:"market"`
	District    string  `json:"district"`
	State       string  `json:"state"`
	MinPrice    float64 `json:"min_price"`
	MaxPrice    float64 `json:"max_price"`
	ModalPrice  float64 `json:"modal_price"`
	ArrivalDate string  `json:"arrival_date"`
}
type GetPriceRequest struct {
	Crop     string `form:"crop"`
	State    string `form:"state"`
	District string `form:"district"`
}
type TrendingPrice struct {
    Crop       string  `db:"crop"`
    Market     string  `db:"market"`
    ModalPrice float64 `db:"modal_price"`
    Change     float64 `db:"change"`
    Trend      string  `db:"trend"`
}
type TrendingPriceResponse struct {
	Crop        string  `json:"crop"`
	Variety     string  `json:"variety,omitempty"`
	Market      string  `json:"market"`
	District    string  `json:"district,omitempty"`
	State       string  `json:"state,omitempty"`
	MinPrice    float64 `json:"min_price"`
	MaxPrice    float64 `json:"max_price"`
	ModalPrice  float64 `json:"modal_price"`
	PriceChange float64 `json:"price_change,omitempty"`
	Trend        string `json:"trend,omitempty"`
	ArrivalDate string  `json:"arrival_date,omitempty"`
}

type PriceComparisonResponse struct {
	Crop          string         `json:"crop"`
	Market        string         `json:"market"`
	CurrentPrice  float64        `json:"current_price"`
	Last7Days     []PriceHistory `json:"last_7_days"`
	Last10Days    []PriceHistory `json:"last_10_days"`
	Last30Days    []PriceHistory `json:"last_30_days"`
}

type PriceHistory struct {
	Date       string  `json:"date"`
	ModalPrice float64 `json:"modal_price"`
	MinPrice   float64 `json:"min_price"`
	MaxPrice   float64 `json:"max_price"`
}