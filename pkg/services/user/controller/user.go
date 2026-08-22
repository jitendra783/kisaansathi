package controller

// import (
// 	"context"
// 	"kisaanSathi/pkg/logger"
// 	"kisaanSathi/pkg/services/user/models"

// 	"go.uber.org/zap"
// )

// func (s *controller) Login(ctx context.Context, request *models.LoginRequest) (*models.LoginResponse, error) {
// 	logger.Log(ctx).Debug("START Login")
// 	defer logger.Log(ctx).Debug("END Login")

// 	result, err := s.registerStore.Login(ctx, request.Email, request.Password)
// 	if err != nil {
// 		logger.Log(ctx).Error("Login failed", zap.Error(err))
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (s *controller) Logout(ctx context.Context, request *models.LogoutRequest) (*models.LogoutResponse, error) {
// 	logger.Log(ctx).Debug("START Logout")
// 	defer logger.Log(ctx).Debug("END Logout")

// 	err := s.registerStore.Logout(ctx, request.LogoutFlag)
// 	if err != nil {
// 		logger.Log(ctx).Error("Logout failed", zap.Error(err))
// 		return nil, err
// 	}
// 	return &models.LogoutResponse{Message: "Logged out successfully"}, nil
// }

// func (s *controller) Register(ctx context.Context, request *models.RegisterRequest) error {
// 	logger.Log(ctx).Debug("START Register")
// 	defer logger.Log(ctx).Debug("END Register")

// 	err := s.registerStore.Register(ctx, request.Mobile, request.Email, request.Password)
// 	if err != nil {
// 		logger.Log(ctx).Error("Error in register", zap.Error(err))
// 		return err
// 	}
// 	return nil
// }

// func (s *controller) GetUserDetails(ctx context.Context, email string) (*models.UserDetails, error) {
// 	logger.Log(ctx).Debug("START Register")
// 	defer logger.Log(ctx).Debug("END Register")

// 	userDetails, err := s.registerStore.GetUserDetails(ctx, email)
// 	if err != nil {
// 		logger.Log(ctx).Error("Error in register", zap.Error(err))
// 		return nil, err
// 	}
// 	return userDetails, nil
// }

// func (s *controller) UserupdateDetails(ctx context.Context, request *models.UpdateUserRequest) (*models.UserDetails, error) {
// 	logger.Log(ctx).Debug("START UserupdateDetails")
// 	defer logger.Log(ctx).Debug("END UserupdateDetails")

// 	// Implementation for updating user details goes here
// 	userDetails, err := s.registerStore.UserupdateDetails(ctx, request.Email, request.Bio)
// 	if err != nil {
// 		logger.Log(ctx).Error("Error in updating user details", zap.Error(err))
// 		return nil, err
// 	}
// 	return userDetails, nil
// }

// func (s *controller) UpdateAvatar(ctx context.Context, email string, avatarURL string) error {
// 	logger.Log(ctx).Debug("START UpdateAvatar")
// 	defer logger.Log(ctx).Debug("END UpdateAvatar")

// 	// Implementation for updating avatar goes here
// 	err := s.registerStore.UpdateAvatar(ctx, email, avatarURL)
// 	if err != nil {
// 		logger.Log(ctx).Error("Error in updating avatar", zap.Error(err))
// 		return err
// 	}
// 	return nil
// }

// // func (s *controller) RefreshToken(ctx context.Context, request *models.DtlsRequest) ([]*models.DtlsResponse, error) {
// // 	logger.Log(ctx).Debug("START RefreshToken")
// // 	defer logger.Log(ctx).Debug("END RefreshToken")

// // 	// Implementation for refreshing token goes here
// // 	data, err := s.registerStore.RefreshToken(ctx, request)
// // 	if err != nil {
// // 		logger.Log(ctx).Error("Error in refreshing token", zap.Error(err))
// // 		return nil, err
// // 	}
// // 	return data, nil
// // }
