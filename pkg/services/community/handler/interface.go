package handler

import (
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/community/controller"
	"kisaanSathi/pkg/services/community/db"

	"github.com/gin-gonic/gin"
)

type communityHandler struct {
	controller controller.CommunityController
}
type CommunityHandler interface {
	GetPosts(ctx *gin.Context)

	CreatePost(ctx *gin.Context)

	UpdatePost(ctx *gin.Context)

	DeletePost(ctx *gin.Context)

	CreateComment(ctx *gin.Context)

	DeleteComment(ctx *gin.Context)

	LikePost(ctx *gin.Context)

	SharePost(ctx *gin.Context)
}

func NewCommunityHandler(ctrl controller.CommunityController) CommunityHandler {
	return &communityHandler{
		controller: ctrl,
	}
}

func NewCommunityController(repo repo.DataObject) controller.CommunityController {
	store := db.NewCommunityStore(repo.Databases.PgDB)
	return controller.NewCommunityController(store)
}
