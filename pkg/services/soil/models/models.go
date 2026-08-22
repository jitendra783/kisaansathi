package models

import "time"

type SoilType struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

type SoilReport struct {
	ID             int64     `db:"id" json:"id"`
	UserID         int64     `db:"user_id" json:"user_id"`
	SoilType       string    `db:"soil_type" json:"soil_type"`
	PH             float64   `db:"ph" json:"ph"`
	Nitrogen       float64   `db:"nitrogen" json:"nitrogen"`
	Phosphorus     float64   `db:"phosphorus" json:"phosphorus"`
	Potassium      float64   `db:"potassium" json:"potassium"`
	OrganicCarbon  float64   `db:"organic_carbon" json:"organic_carbon"`
	ElectricalEC   float64   `db:"electrical_ec" json:"electrical_ec"`
	Moisture       float64   `db:"moisture" json:"moisture"`
	Status         string    `db:"status" json:"status"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}

type SoilRecommendation struct {
	ID             int64  `db:"id" json:"id"`
	Crop           string `db:"crop" json:"crop"`
	SoilType       string `db:"soil_type" json:"soil_type"`
	Recommendation string `db:"recommendation" json:"recommendation"`
	Fertilizer     string `db:"fertilizer" json:"fertilizer"`
	Irrigation     string `db:"irrigation" json:"irrigation"`
}