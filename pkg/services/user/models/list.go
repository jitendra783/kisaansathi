package models

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LogoutRequest struct {
	LogoutFlag string `json:"logoutFlag" binding:"required,oneof=Y N"`
}
type LoginResponse struct {
	Token   string `json:"FML_COMP_CD,omitempty"`
	XLength string `json:"FML_LM_FLG,omitempty"`
}

type MfListDetails struct {
	MFCompCd     string `gorm:"column:MF_COMP_CD1"`
	MFCompTiFlag string `gorm:"column:MF_COMP_TI_FLG1"`
	MFCompName   string `gorm:"column:MF_COMP_NAME1"`
}

type RegisterRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
	Email     string `json:"email" binding:"required"`
	ZipCode   string `json:"zipcode" binding:"required"`
	State     string `json:"state" binding:"required"`
	District  string `json:"district" binding:"required"`
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
type SipprotectAmcListRequest struct {
	FML_MATCH_ACCNT string `json:"FML_MATCH_ACCNT" binding:"matchaccount"`
	FML_TRNSCTN_FLW string `json:"FML_TRNSCTN_FLW" binding:"omitempty,oneof=Y N"`
	FML_PRDCT_TYP   string `json:"FML_PRDCT_TYP" binding:"omitempty,oneof=Y N"`
	FML_RQST_TYP    string `json:"FML_RQST_TYP" binding:"omitempty,oneof=Y N"`
}

type MfListResponse struct {
	FML_COMP_CD   string `json:"FML_COMP_CD,omitempty"`
	FML_COMP_NAME string `json:"FML_COMP_NAME,omitempty"`
}
