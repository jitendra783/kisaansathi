package models


// ==============================
// Create Banner
// ==============================

type CreateBannerRequest struct {
	Page         string `json:"page" binding:"required"`
	Title        string `json:"title" binding:"required"`
	Subtitle     string `json:"subtitle"`
	ImageURL     string `json:"imageUrl" binding:"required"`
	ActionURL    string `json:"actionUrl"`
	DisplayOrder int    `json:"displayOrder"`
	IsActive     bool   `json:"isActive"`
}

// ==============================
// Update Banner
// ==============================

type UpdateBannerRequest struct {
	ID           int64  `json:"id" binding:"required"`
	Page         string `json:"page"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	ImageURL     string `json:"imageUrl"`
	ActionURL    string `json:"actionUrl"`
	DisplayOrder int    `json:"displayOrder"`
	IsActive     bool   `json:"isActive"`
}

// ==============================
// Response
// ==============================

type BannerResponse struct {
	ID           int64  `json:"id"`
	Page         string `json:"page"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	ImageURL     string `json:"imageUrl"`
	ActionURL    string `json:"actionUrl"`
	DisplayOrder int    `json:"displayOrder"`
	IsActive     bool   `json:"isActive"`
}

// ==============================
// List Response
// ==============================

type GetBannersResponse struct {
	Banners []*BannerResponse `json:"banners"`
}