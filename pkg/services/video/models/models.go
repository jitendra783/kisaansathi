package models

type Video struct {
	ID          int64  `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Thumbnail   string `db:"thumbnail" json:"thumbnail"`
	VideoURL    string `db:"video_url" json:"video_url"`
	Category    string `db:"category" json:"category"`
	Duration    string `db:"duration" json:"duration"`
	Language    string `db:"language" json:"language"`
	CreatedAt   string `db:"created_at" json:"created_at"`
}

type VideoResponse struct {
	Data *Video `json:"data"`
}

type VideosResponse struct {
	Data []Video `json:"data"`
}