package models

type FAQ struct {
	ID       int64  `db:"id" json:"id"`
	Question string `db:"question" json:"question"`
	Answer   string `db:"answer" json:"answer"`
}

type FAQResponse struct {
	Data []FAQ `json:"data"`
}

type FAQDetailResponse struct {
	Data FAQ `json:"data"`
}
