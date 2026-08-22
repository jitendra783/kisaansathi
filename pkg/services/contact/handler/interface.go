package handler

import (
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/contact/controller"
	"kisaanSathi/pkg/services/contact/db"

	"github.com/gin-gonic/gin"
)

type contactHandler struct {
	controller controller.ContactController
}
type ContactHandler interface {
	ContactUs(ctx *gin.Context)
	GetContactInfo(ctx *gin.Context)
}

func NewContacthandler(controller controller.ContactController) ContactHandler {
	return &contactHandler{
		controller: controller,
	}
}
func NewContactController(repo repo.DataObject) controller.ContactController {
	store := db.NewContactStore(repo.Databases.PgDB)
	return controller.NewContactController(store)
}
