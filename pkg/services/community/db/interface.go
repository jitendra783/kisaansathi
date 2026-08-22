package db

import (
	"kisaanSathi/pkg/services/community/models"

	"github.com/jmoiron/sqlx"
)

type communityStore struct {
	db *sqlx.DB
}

type CommunityStore interface {
	GetPosts() ([]models.Post, error)

	CreatePost(req *models.CreatePostRequest) error

	UpdatePost(id int64, req *models.UpdatePostRequest) error

	DeletePost(id int64) error

	CreateComment(req *models.CreateCommentRequest) error

	DeleteComment(id int64) error

	LikePost(req *models.LikeRequest) error

	SharePost(req *models.ShareRequest) error
}

func NewCommunityStore(db *sqlx.DB) CommunityStore {
	return &communityStore{
		db: db,
	}
}
