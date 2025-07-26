package mandi

import (
	"kisaanSathi/pkg/repo"

	"github.com/gin-gonic/gin"
)

type mandiHandler struct {
}
type MandiHandler interface {
	GetMandiBhav(c *gin.Context)
}

func NewMandiHandler(repo repo.DataObject) MandiHandler {
	return &mandiHandler{}
}
