package controller

import (
	"context"

	"kisaanSathi/pkg/services/user/db"
	"kisaanSathi/pkg/services/user/models"
	"kisaanSathi/pkg/storage"

	"mime/multipart"
)

type controller struct {
	store   db.UserStore
	storage storage.Storage
}

type UserController interface {
	// Authentication
	Register(ctx context.Context, req *models.RegisterRequest) error
	Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error)
	Logout(ctx context.Context, req *models.LogoutRequest) (*models.LogoutResponse, error)

	// Profile
	GetUserDetails(ctx context.Context, email string) (*models.UserDetailsResponse, error)
	UpdateUserDetails(ctx context.Context, req *models.UpdateUserRequest) (*models.UpdateUserResponse, error)
	UpdateAvatar(ctx context.Context, email string, file *multipart.FileHeader) (*models.UpdateAvatarResponse, error)

	// Password
	ChangePassword(ctx context.Context, req *models.ChangePasswordRequest) (*models.ChangePasswordResponse, error)
	ForgotPassword(ctx context.Context, req *models.ForgotPasswordRequest) (*models.ForgotPasswordResponse, error)
	ResetPassword(ctx context.Context, req *models.ResetPasswordRequest) (*models.ResetPasswordResponse, error)

	// Token
	RefreshToken(ctx context.Context, req *models.RefreshTokenRequest) (*models.RefreshTokenResponse, error)

	// User
	DeleteUser(ctx context.Context, email string) (*models.DeleteUserResponse, error)
}

func NewUserController(store db.UserStore, storage storage.Storage) UserController {
	return &controller{
		store:   store,
		storage: storage,
	}
}
