package controller

import (
	"context"
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/user/models"
)

func (s *controller) Login(ctx context.Context, request *models.LoginRequest) (data []*models.LoginResponse, err error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	var matchAccount = request.Email
	result, err := s.registerStore.Login(ctx, matchAccount)
	return result, err

}
func (s *controller) Logout(ctx context.Context, request *models.LogoutRequest) (data []*models.LoginResponse, err error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")
	result, err := s.registerStore.Login(ctx, request.LogoutFlag)
	return result, err
}

func (s *controller) Register(ctx context.Context, request *models.RegisterRequest) (data []*models.RegisterResponse, err error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")
	result, err := s.registerStore.Register(ctx, "", "", "")
	return result, err
}

// func (s *controller) RefreshToken(ctx context.Context, request *models.RefreshTokenRequest) (data []*models.Refre, err error) {
// 	logger.Log(ctx).Debug("START")
// 	defer logger.Log(ctx).Debug("END")
// 	// var comp_cd = request.FML_COMP_CD
// 	// var sch_cd = request.FML_MF_SCH_CD
// 	// var ld_type = request.FML_MF_LD_TYPE
// 	logger.Log(ctx).Debug("company code:", zap.String("value", request.FML_COMP_CD))
// 	logger.Log(ctx).Debug("scheme code:", zap.String("value", request.FML_MF_SCH_CD))
// 	logger.Log(ctx).Debug("ld type:", zap.String("value", request.FML_MF_LD_TYPE))
// 	logger.Log(ctx).Debug("Check input  details")

// 	//result, err := s.db.GetDtlsDetails(ctx, comp_cd, sch_cd, ld_type)
// 	return data, err
// }
