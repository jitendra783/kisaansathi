package models

import "time"

type FeedbackRequest struct {
	Name       string `db:"name" json:"name"`
	Phone      string `db:"phone" json:"phone"`
	Email      string `db:"email" json:"email"`
	Rating     int    `db:"rating" json:"rating"`
	Category   string `db:"category" json:"category"`
	Message    string `db:"message" json:"message"`
	AppVersion string `db:"app_version" json:"app_version"`
}

type FeedbackDBResult struct {
	ID         int64     `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Phone      string    `db:"phone" json:"phone"`
	Email      string    `db:"email" json:"email"`
	Rating     int       `db:"rating" json:"rating"`
	Category   string    `db:"category" json:"category"`
	Message    string    `db:"message" json:"message"`
	AppVersion string    `db:"app_version" json:"app_version"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

type Feedback struct {
	ID         int64     `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Phone      string    `json:"phone" db:"phone"`
	Email      string    `json:"email" db:"email"`
	Rating     int       `json:"rating" db:"rating"`
	Category   string    `json:"category" db:"category"`
	Message    string    `json:"message" db:"message"`
	AppVersion string    `json:"app_version" db:"app_version"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type FeedbackListResponse struct {
	Data []Feedback `json:"data"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
