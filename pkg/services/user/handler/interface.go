package handler

import (
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/user/controller"
	"kisaanSathi/pkg/services/user/db"

	"github.com/gin-gonic/gin"
)

type handler struct {
	controller controller.RegisterController
}

type RegisterHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	//RefreshToken(c *gin.Context)
	Logout(c *gin.Context)
}

func NewRegisterHandler(controller controller.RegisterController) RegisterHandler {
	return &handler{
		controller: controller,
	}
}
func RegisterController(repo repo.DataObject) controller.RegisterController {
	store := db.NewDBObject(repo.Databases.PgDB)
	return controller.NewRegisterController(store)
}
