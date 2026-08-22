package handler

import (
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/home/controller"
	"kisaanSathi/pkg/services/home/db"

	"github.com/gin-gonic/gin"
)

type homeHandler struct {
	controller controller.HomeController
}
type HomeHandler interface {
	GetHome(*gin.Context)
	GetDashboardData(*gin.Context)
}

func NewHomeHandler(controller controller.HomeController) HomeHandler {
	return &homeHandler{
		controller: controller,
	}
}
func NewHomeController(repo repo.DataObject) controller.HomeController {
	store := db.NewHomeStore(repo.Databases.PgDB)
	return controller.NewHomeController(store)
}
