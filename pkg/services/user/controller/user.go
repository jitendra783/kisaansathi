package controller

import (
	"context"
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/user/models"

	"go.uber.org/zap"
)

func (s *controller) Login(ctx context.Context, request *models.LoginRequest) (*models.LoginResponse, error) {
	logger.Log(ctx).Debug("START Login")
	defer logger.Log(ctx).Debug("END Login")

	result, err := s.registerStore.Login(ctx, request.Email, request.Password)
	if err != nil {
		logger.Log(ctx).Error("Login failed", zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (s *controller) Logout(ctx context.Context, request *models.LogoutRequest) (*models.LogoutResponse, error) {
	logger.Log(ctx).Debug("START Logout")
	defer logger.Log(ctx).Debug("END Logout")

	err := s.registerStore.Logout(ctx, request.LogoutFlag)
	if err != nil {
		logger.Log(ctx).Error("Logout failed", zap.Error(err))
		return nil, err
	}
	return &models.LogoutResponse{Message: "Logged out successfully"}, nil
}

func (s *controller) Register(ctx context.Context, request *models.RegisterRequest) error {
	logger.Log(ctx).Debug("START Register")
	defer logger.Log(ctx).Debug("END Register")

	err := s.registerStore.Register(ctx, request.Mobile, request.Email, request.Password)
	if err != nil {
		logger.Log(ctx).Error("Error in register", zap.Error(err))
		return err
	}
	return nil
}
