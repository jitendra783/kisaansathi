package handler

import (
	"kisaanSathi/pkg/repo"
	repo_session "kisaanSathi/pkg/services/session/controller"

	"github.com/gin-gonic/gin"
)

type sessionObject struct {
	repo repo_session.RepoLayer
}

type SessionGroup interface {
	Validate(c *gin.Context)
}

func NewSessionGroup(repo repo.DataObject) SessionGroup {
	return &sessionObject{
		repo: repo_session.NewRepoLayerObject(repo),
	}
}
