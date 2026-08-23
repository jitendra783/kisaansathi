package controller

import (
	"kisaanSathi/pkg/services/community/db"
	"kisaanSathi/pkg/services/community/models"
)

type communityController struct {
	store db.CommunityStore
}
type CommunityController interface {
	GetPosts() (*models.PostListResponse, error)

	CreatePost(req *models.CreatePostRequest) (*models.CommonResponse, error)

	UpdatePost(id int64, req *models.UpdatePostRequest) (*models.CommonResponse, error)

	DeletePost(id int64) (*models.CommonResponse, error)

	CreateComment(req *models.CreateCommentRequest) (*models.CommonResponse, error)

	DeleteComment(id int64) (*models.CommonResponse, error)

	LikePost(req *models.LikeRequest) (*models.CommonResponse, error)

	SharePost(req *models.ShareRequest) (*models.CommonResponse, error)
}

func NewCommunityController(store db.CommunityStore) CommunityController {
	return &communityController{
		store: store,
	}
}
