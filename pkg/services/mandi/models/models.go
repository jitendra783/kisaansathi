package models

import (
	"database/sql"
)

type MandiPrice struct {
	ID          sql.NullString   `db:"id"`
	Crop        sql.NullString  `db:"crop"`
	Variety     sql.NullString  `db:"variety"`
	Market      sql.NullString  `db:"market"`
	District    sql.NullString  `db:"district"`
	State       sql.NullString  `db:"state"`
	MinPrice    sql.NullFloat64 `db:"min_price"`
	MaxPrice    sql.NullFloat64 `db:"max_price"`
	ModalPrice  sql.NullFloat64 `db:"modal_price"`
	ArrivalDate sql.NullTime    `db:"arrival_date"`
}

type MandiPriceResponse struct {
	ID          int   `json:"id"`
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
	Crop       sql.NullString  `db:"crop"`
	Market     sql.NullString  `db:"market"`
	ModalPrice sql.NullFloat64 `db:"modal_price"`
	Change     sql.NullFloat64 `db:"change"`
	Trend      sql.NullString  `db:"trend"`
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
	Trend       string  `json:"trend,omitempty"`
	ArrivalDate string  `json:"arrival_date,omitempty"`
}

type PriceComparisonResponse struct {
	Crop         string         `json:"crop"`
	Market       string         `json:"market"`
	CurrentPrice float64        `json:"current_price"`
	Last7Days    []PriceHistory `json:"last_7_days"`
	Last10Days   []PriceHistory `json:"last_10_days"`
	Last30Days   []PriceHistory `json:"last_30_days"`
}

type PriceHistory struct {
	Date       string  `json:"date"`
	ModalPrice float64 `json:"modal_price"`
	MinPrice   float64 `json:"min_price"`
	MaxPrice   float64 `json:"max_price"`
}
