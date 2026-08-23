package models

type Banner struct {
	ID       int    `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	ImageURL string `db:"image_url" json:"image_url"`
	Redirect string `db:"redirect_url" json:"redirect_url"`
	OrderNo  int    `db:"order_no" json:"order_no"`
}

type HomeCard struct {
	ID       int    `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	Icon     string `db:"icon" json:"icon"`
	Redirect string `db:"redirect_url" json:"redirect_url"`
	OrderNo  int    `db:"order_no" json:"order_no"`
}

type HomeResponse struct {
	Banners []Banner   `json:"banners"`
	Cards   []HomeCard `json:"cards"`
}
