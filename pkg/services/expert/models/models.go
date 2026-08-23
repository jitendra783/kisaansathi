package models

// =========================
// Expert
// =========================

type Expert struct {
	ID              int64   `db:"id" json:"id"`
	Name            string  `db:"name" json:"name"`
	Specialization  string  `db:"specialization" json:"specialization"`
	Experience      int     `db:"experience" json:"experience"`
	Language        string  `db:"language" json:"language"`
	Rating          float64 `db:"rating" json:"rating"`
	Bio             string  `db:"bio" json:"bio"`
	ConsultationFee float64 `db:"consultation_fee" json:"consultation_fee"`
	Phone           string  `db:"phone" json:"phone"`
	ProfileImage    string  `db:"profile_image" json:"profile_image"`
	Available       bool    `db:"available" json:"available"`
}

// =========================
// Get Experts Response
// =========================

type GetExpertsResponse struct {
	Data []Expert `json:"data"`
}

// =========================
// Get Expert Response
// =========================

type GetExpertResponse struct {
	Data *Expert `json:"data"`
}

// =========================
// Book Consultation Request
// =========================

type BookConsultationRequest struct {
	UserID   int64  `json:"user_id" binding:"required"`
	ExpertID int64  `json:"expert_id" binding:"required"`
	Date     string `json:"date" binding:"required"`
	Time     string `json:"time" binding:"required"`
	Mode     string `json:"mode" binding:"required"`
	Problem  string `json:"problem" binding:"required"`
}

// =========================
// Book Consultation Response
// =========================

type BookConsultationResponse struct {
	Message string `json:"message"`
}

// =========================
// Consultation History
// =========================

type ConsultationHistory struct {
	BookingID  int64  `db:"booking_id" json:"booking_id"`
	ExpertName string `db:"expert_name" json:"expert_name"`
	Date       string `db:"date" json:"date"`
	Time       string `db:"time" json:"time"`
	Mode       string `db:"mode" json:"mode"`
	Status     string `db:"status" json:"status"`
}

// =========================
// Consultation History Response
// =========================

type ConsultationHistoryResponse struct {
	Data []ConsultationHistory `json:"data"`
}
