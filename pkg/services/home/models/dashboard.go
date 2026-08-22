package models

type DashboardResponse struct {
	Weather interface{} `json:"weather"`
	Mandi   interface{} `json:"mandi"`
	Feeds   interface{} `json:"feeds"`
	User    interface{} `json:"user"`
}
