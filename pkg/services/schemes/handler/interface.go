package handler

import (
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/schemes/controller"
	"kisaanSathi/pkg/services/schemes/db"

	"github.com/gin-gonic/gin"
)

type schemeHandler struct {
	controller controller.SchemeController
}

type SchemeHandler interface {
	GetSchemes(ctx *gin.Context)
	GetScheme(ctx *gin.Context)
	GetEligibleSchemes(ctx *gin.Context)
	GetStateSchemes(ctx *gin.Context)
}

func NewSchemeHandler(controller controller.SchemeController) SchemeHandler {
	return &schemeHandler{
		controller: controller,
	}
}
func NewSchemeController(repo repo.DataObject) controller.SchemeController {
	store := db.NewSchemeStore(repo.Databases.PgDB)
	return controller.NewSchemeController(store)
}
