package models

import "time"

// Define structs to represent database table rows
type UsmSssnMngr struct {
	UsmUsrID          string
	UsmSssnID         int
	UsmSssnEndDtFlg   string
	UsmSssnLstAccsDt  time.Time
	UsmSupUsrTyp      string
	UsmIPID           string
	UsmSssnTermntdFlg string
	UsmLoginSource    string
}
