package models

type UserInfo struct {
	PersonalInfo
	ExtraPersonalInfo
	CustomerInfo
}

type ExtraInfo struct {
	ExtraPersonalInfo
	CustomerInfo
}

type PersonalInfo struct {
	FullName string `json:"FML_USR_USR_NM,omitempty"`
	Email    string `json:"FML_URQ_RQST_DTLS,omitempty"`
	MobileNo string `json:"FML_MDC_CRDT_NMBR,omitempty"`
	Gender   string `json:"FML_MMD_IND,omitempty"`
}

type CustomerInfo struct {
	PanNo        string `json:"FML_DEALER_CD,omitempty"`
	IsRICustomer string `json:"FML_PASSWORD_TYPE,omitempty"`
}

type ExtraPersonalInfo struct {
	FirstName     string `json:"FML_USR_FRST_NM,omitempty"`
	MiddleName    string `json:"FML_USR_MDDL_NM,omitempty"`
	LastName      string `json:"FML_USR_LST_NM,omitempty"`
	MatchAccount  string `json:"FML_MATCH_ACCNT,omitempty"`
	ResidentialNo string `json:"FML_USR_PHN_RNMBRS,omitempty"`
	OfficialNo    string `json:"FML_USR_PHN_ONMBRS,omitempty"`
	DateOfBirth   string `json:"FML_USR_DT_BRTH,omitempty"`
	EmailAddress  string `json:"FML_USR_EMAIL_ADDRSS,omitempty"`
}

type PrivacyInfo struct {
	PrivacyOpted  bool
	ModifyAllowed string
}
