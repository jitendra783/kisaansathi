package models

type CreateSoilTestRequest struct {
	UserID        int64   `json:"user_id" binding:"required"`
	SoilType      string  `json:"soil_type" binding:"required"`
	PH            float64 `json:"ph"`
	Nitrogen      float64 `json:"nitrogen"`
	Phosphorus    float64 `json:"phosphorus"`
	Potassium     float64 `json:"potassium"`
	OrganicCarbon float64 `json:"organic_carbon"`
	ElectricalEC  float64 `json:"electrical_ec"`
	Moisture      float64 `json:"moisture"`
}
type SoilTypesResponse struct {
	Data []SoilType `json:"data"`
}

type SoilReportResponse struct {
	Data *SoilReport `json:"data"`
}

type SoilRecommendationResponse struct {
	Data []SoilRecommendation `json:"data"`
}

type CommonResponse struct {
	Message string `json:"message"`
}
