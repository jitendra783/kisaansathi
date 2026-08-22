package models

type Scheme struct {
	ID              int64  `db:"id" json:"id"`
	Name            string `db:"name" json:"name"`
	Description     string `db:"description" json:"description"`
	Category        string `db:"category" json:"category"`
	State           string `db:"state" json:"state"`
	Eligibility     string `db:"eligibility" json:"eligibility"`
	Benefits        string `db:"benefits" json:"benefits"`
	OfficialURL     string `db:"official_url" json:"official_url"`
	ApplicationMode string `db:"application_mode" json:"application_mode"`
	Status          string `db:"status" json:"status"`
}

type SchemeDetailResponse struct {
	Scheme *Scheme `json:"scheme"`
}

type EligibleSchemeRequest struct {
	State      string `form:"state"`
	Category   string `form:"category"`
	FarmerType string `form:"farmer_type"`
}
type StateSchemeRequest struct {
	State string `form:"state"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
}

type SchemeResponse struct {
	Schemes    []Scheme   `json:"schemes"`
	Pagination Pagination `json:"pagination"`
}
