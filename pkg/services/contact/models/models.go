package models

type ContactRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone"`
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
}
type ContactResponse struct {
	Message string `json:"message"`
}


type ContactInfo struct {
	Email     string `db:"email" json:"email"`
	Phone     string `db:"phone" json:"phone"`
	Address   string `db:"address" json:"address"`
	Facebook  string `db:"facebook" json:"facebook"`
	Instagram string `db:"instagram" json:"instagram"`
	Twitter   string `db:"twitter" json:"twitter"`
	LinkedIn  string `db:"linkedin" json:"linkedin"`
}

type ContactInfoResponse struct {
	Data *ContactInfo `json:"data"`
}