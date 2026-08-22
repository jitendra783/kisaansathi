package models

type Banner struct {
	ID           int64  `db:"id" json:"id"`
	Page         string `db:"page" json:"page"`
	Title        string `db:"title" json:"title"`
	Subtitle     string `db:"subtitle" json:"subtitle"`
	ImageURL     string `db:"image_url" json:"image_url"`
	ActionURL    string `db:"action_url" json:"action_url"`
	DisplayOrder int    `db:"display_order" json:"display_order"`
	IsActive     bool   `db:"is_active" json:"is_active"`
}
