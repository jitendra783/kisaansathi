package models

import "time"

type Crop struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Season      string    `db:"season" json:"season"`
	SoilType    string    `db:"soil_type" json:"soil_type"`
	Duration    int       `db:"duration" json:"duration"`
	ImageURL    string    `db:"image_url" json:"image_url"`
	Description string    `db:"description" json:"description"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type RecommendedCrop struct {
	Crop       string `db:"crop" json:"crop"`
	Confidence int    `db:"confidence" json:"confidence"`
}

type CreateCropRequest struct {
	Name        string `json:"name" binding:"required"`
	Season      string `json:"season" binding:"required"`
	SoilType    string `json:"soil_type" binding:"required"`
	Duration    int    `json:"duration" binding:"required"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

type UpdateCropRequest struct {
	Name        string `json:"name" binding:"required"`
	Season      string `json:"season" binding:"required"`
	SoilType    string `json:"soil_type" binding:"required"`
	Duration    int    `json:"duration" binding:"required"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

type CropResponse struct {
	Data *Crop `json:"data"`
}

type CropListResponse struct {
	Data []Crop `json:"data"`
}

type RecommendedCropResponse struct {
	Data []RecommendedCrop `json:"data"`
}
