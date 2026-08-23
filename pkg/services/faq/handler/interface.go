package handler

import (
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/faq/controller"
	"kisaanSathi/pkg/services/faq/db"

	"github.com/gin-gonic/gin"
)

type faqHandler struct {
	controller controller.FAQController
}
type FAQHandler interface {
	GetFAQs(ctx *gin.Context)
	GetFAQByID(ctx *gin.Context)
	GetFAQsByCategory(ctx *gin.Context)
}

func NewFAQHandler(controller controller.FAQController) FAQHandler {
	return &faqHandler{controller: controller}
}
func NewFAQController(repo repo.DataObject) controller.FAQController {
	store := db.NewFAQStore(repo.Databases.PgDB)
	return controller.NewFAQController(store)
}
