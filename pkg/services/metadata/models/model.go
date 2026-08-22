package models

type MetadataDB struct {
	ID        int64  `db:"id"`
	Key       string `db:"key"`
	Value     string `db:"value"`
	Type      string `db:"type"`
	IsActive  bool   `db:"is_active"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
