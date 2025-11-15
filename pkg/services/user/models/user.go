package models

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LogoutRequest struct {
	LogoutFlag string `json:"logoutFlag" binding:"required,oneof=Y N"`
}
type LoginResponse struct {
	Token string `json:"token"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}
type RegisterRequest struct {
	FirstName string `json:"firstName" binding:"required" error:"First name is required"`
	LastName  string `json:"lastName" binding:"omitempty" error:"Last name is required"`
	Address   string `json:"address" binding:"omitempty" error:"Address is required"`
	Mobile    string `json:"mobile" binding:"required" error:"Mobile number is required"`
	Password  string `json:"password" binding:"required" error:"Password is required"`
	//ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password" error:"Password and Confirm Password must be same"`
	Email    string `json:"email" binding:"required" error:"Email is required"`
	ZipCode  string `json:"zipcode" binding:"omitempty" error:"Zipcode is required"`
	State    string `json:"state" binding:"omitempty" error:"State is required"`
	District string `json:"district" binding:"omitempty" error:"District is required"`
}

type RegisterResponse struct {
	FML_MF_LD_CAT         string `json:"FML_MF_LD_CAT,omitempty"`
	FML_MF_LD_START_RANGE string `json:"FML_MF_LD_START_RANGE,omitempty"`
	FML_MF_LD_END_RANGE   string `json:"FML_MF_LD_END_RANGE,omitempty"`
	FML_MF_LD_PERCENTAGE  string `json:"FML_MF_LD_PERCENTAGE,omitempty"`
	FML_MF_LD_MIN_LOAD    string `json:"FML_MF_LD_MIN_LOAD,omitempty"`
	FML_MF_LD_MAX_LOAD    string `json:"FML_MF_LD_MAX_LOAD,omitempty"`
	FML_MF_LD_REMARKS     string `json:"FML_MF_LD_REMARKS,omitempty"`
}

type RefreshTokenRequest struct {
	Email              string `json:"FML_MATCH_ACCNT" binding:"required,matchaccount" error:"Provide valid Match account"`
	FML_NOMINATION_FLG string `json:"FML_NOMINATION_FLG" binding:"omitempty,oneof=Y N"`
	FML_RQST_TYP       string `json:"FML_RQST_TYP" binding:"omitempty,oneof=Y N"`
}

type User struct {
	ID       int64  `gorm:"primaryKey" json:"-"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Language string `json:"language"`
	SoilType string `json:"soil_type"`
	District string `json:"district"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

