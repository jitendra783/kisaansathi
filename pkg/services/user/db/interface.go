package db

import (
	"context"
	"kisaanSathi/pkg/services/user/models"

	"gorm.io/gorm"
)

type registerStore struct {
	store *gorm.DB
}

type RegisterStore interface {
	Login(context.Context, string) ([]*models.LoginResponse, error)
	Logout(context.Context, string, string, string) ([]*models.LoginResponse, error)
	Register(context.Context, string, string, string) ([]*models.RegisterResponse, error)
}

func NewDBObject(store *gorm.DB) RegisterStore {
	return &registerStore{store: store}
}
