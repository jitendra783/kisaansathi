package models

type MetadataResponse struct {
	Stats      StatsResponse      `json:"stats"`
	Categories []CategoryResponse `json:"categories"`
	Languages  []string           `json:"languages"`
	CropTypes  []string           `json:"cropTypes"`
}

type StatsResponse struct {
	Farmers  string `json:"farmers"`
	Experts  string `json:"experts"`
	Districts string `json:"districts"`
	Support  string `json:"support"`
}

type CategoryResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Order int    `json:"order"`
}