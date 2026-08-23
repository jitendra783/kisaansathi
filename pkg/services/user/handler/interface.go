package handler

import (
	"kisaanSathi/pkg/repo"
	"kisaanSathi/pkg/services/user/controller"
	"kisaanSathi/pkg/services/user/db"
	"kisaanSathi/pkg/storage"

	"github.com/gin-gonic/gin"
)

type handler struct {
	controller controller.UserController
}

type UserHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Register(c *gin.Context)

	GetUserDetails(c *gin.Context)
	UpdateUserDetails(c *gin.Context)
	UpdateAvatar(c *gin.Context)

	RefreshToken(c *gin.Context)
	ChangePassword(c *gin.Context)
	ForgotPassword(c *gin.Context)
	ResetPassword(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func NewUserHandler(controller controller.UserController) UserHandler {
	return &handler{
		controller: controller,
	}
}

func NewUserController(repo repo.DataObject) controller.UserController {
	store := db.NewDBObject(repo.Databases.PgDB)
	storage := storage.NewLocalStorage()
	return controller.NewUserController(store, storage)
}
