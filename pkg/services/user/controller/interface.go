package controller

import (
	"context"
	"kisaanSathi/pkg/services/user/db"
	"kisaanSathi/pkg/services/user/models"
)

type controller struct {
	registerStore db.RegisterStore
}

type RegisterController interface {
	Login(ctx context.Context, request *models.LoginRequest) (data []*models.LoginResponse, err error)
	Logout(ctx context.Context, request *models.LogoutRequest) (data []*models.LoginResponse, err error)
	Register(ctx context.Context, request *models.RegisterRequest) (data []*models.RegisterResponse, err error)
	//RefreshToken(ctx context.Context, request *models.DtlsRequest) (data []*models.DtlsResponse, err error)
}

func NewRegisterController(registerStore db.RegisterStore) RegisterController {
	return &controller{
		registerStore: registerStore,
	}
}
