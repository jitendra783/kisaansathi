package feeds

import (
	"kisaanSathi/pkg/repo"

	"github.com/gin-gonic/gin"
)

type feedsHandler struct {
}
type FeedsHandler interface {
	GetFeeds(c *gin.Context)
}

func NewFeedsHandler(repo repo.DataObject) FeedsHandler {
	return &feedsHandler{}
}
