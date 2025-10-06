package models

type MandiPrice struct {
	Market      string  `json:"market"`
	Commodity   string  `json:"commodity"`
	ModalPrice  float64 `json:"modal_price"`
	MinPrice    float64 `json:"min_price"`
	MaxPrice    float64 `json:"max_price"`
	Date        string  `json:"date"`
}