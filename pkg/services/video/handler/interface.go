package handler

import (
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/video/controller"
	"kisaanSathi/pkg/services/video/db"

	"github.com/gin-gonic/gin"
)

type videoHandler struct {
	controller controller.VideoController
}
type VideoHandler interface {
	GetVideos(ctx *gin.Context)
	GetVideo(ctx *gin.Context)
	GetVideosByCategory(ctx *gin.Context)
}

func NewVideoHandler(controller controller.VideoController) VideoHandler {
	return &videoHandler{controller: controller}
}
func NewVideoController(repo repo.DataObject) controller.VideoController {
	store := db.NewVideoStore(repo.Databases.PgDB)
	return controller.NewVideoController(store)
}
