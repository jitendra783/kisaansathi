package db

import (
	"context"
	"errors"
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/user/models"

	"go.uber.org/zap"
)

func (g *registerStore) Login(c context.Context, matchAccount string) ([]*models.LoginResponse, error) {
	logger.Log(c).Debug("START")
	logger.Log(c).Debug("END")

	var items []*models.LoginResponse // Assuming scheme.Item is the response type

	var mfLists []models.MfListDetails
	// Perform the query
	// i_flg_nri := utils.Strings(matchAccount).IsNriUser()
	// query := `	SELECT  mf_comp_cd AS "MF_COMP_CD1",
	//                   	mf_comp_name AS "MF_COMP_NAME1",
	//                   	mf_comp_ti_flg AS "MF_COMP_TI_FLG1"
	// 			FROM  	mf_companies
	// 			WHERE 	nvl(MF_COMP_STATUS_FLG,'N') = 'Y'
	// 			AND		MF_COMP_NRI_FLG = decode(?, 'N', MF_COMP_NRI_FLG, 'Y')
	// 			ORDER BY mf_comp_name`
	// // Scan the results into the items slice
	// err := g.store.WithContext(c).Raw(query, i_flg_nri).Scan(&mfLists).Error
	// if err != nil {
	// 	logger.Log(c).Error("Error executing query", zap.Error(err))
	// 	return nil, err
	// }

	// logger.Log(c).Debug("Query:", zap.Any("query", query))

	for _, mf := range mfLists {
		items = append(items, &models.LoginResponse{
			Token:   mf.MFCompCd,
			XLength: mf.MFCompName,
		})
	}
	logger.Log(c).Debug("Result:", zap.Any("result:", items))
	if items == nil {
		return items, errors.New("data not found")
	}
	return items, nil
}
func (g *registerStore) Logout(c context.Context, matchAccount string, nominationFlg string, rqstType string) ([]*models.LoginResponse, error) {
	logger.Log(c).Debug("START")
	logger.Log(c).Debug("END")
	var logoutResponse []*models.LoginResponse
	return logoutResponse, nil
}

func (g *registerStore) Register(c context.Context, matchAccount string, sipProductFlg string, sipTransactionFlg string) ([]*models.RegisterResponse, error) {
	logger.Log(c).Debug("START")
	logger.Log(c).Debug("END")

	var items []*models.RegisterResponse

	logger.Log(c).Debug("Result:", zap.Any("result:", items))
	if items == nil {
		return items, errors.New("data not found")
	}
	return items, nil
}
