package models

type GeoResponse struct {
	Results []struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Name      string  `json:"name"`
		Country   string  `json:"country"`
		Timezone  string  `json:"timezone"`
	} `json:"results"`
}


type MandiPrice struct {
	Market      string  `json:"market"`
	Commodity   string  `json:"commodity"`
	ModalPrice  float64 `json:"modal_price"`
	MinPrice    float64 `json:"min_price"`
	MaxPrice    float64 `json:"max_price"`
	Date        string  `json:"date"`
}