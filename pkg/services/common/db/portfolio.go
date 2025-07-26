package db

import (
	"context"
	"database/sql"
	"errors"
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/common/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type portfolioDB struct {
	oracle *gorm.DB
}

type PortfolioStore interface {
	IsTransactionExistsForScheme(ctx context.Context, compCode int, schemeCode, matchAccount string) (bool, error)
	GetDPDetails(ctx context.Context, compCode int, schemeCode, folioNo, matchAccount string) (*models.DpDetails, error)
	GetDematUnblockedUnits(ctx context.Context, matchAccount, dpID, dpClientID, stockCode, isinNumber string) (*models.UnblockedUnits, error)
	GetUnblockedUnits(ctx context.Context, compCode int, schemeCode, folioNo, matchAccount string) (*models.UnblockedUnits, error)
}

func NewPortfolioStore(oracle *gorm.DB) PortfolioStore {
	return &portfolioDB{oracle: oracle}
}

func (p portfolioDB) IsTransactionExistsForScheme(ctx context.Context, compCode int, schemeCode, matchAccount string) (bool, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	var transCount uint16
	query := p.oracle.WithContext(ctx).
		Select(`COUNT(*) AS "transCount"`).
		Table(`MF_UNIT_BAL`).
		Where(`MF_UNB_MATCH_ACC = ?`, matchAccount).
		Where(`MF_UNB_COMP_CD = ?`, compCode).
		Where(`MF_UNB_SCH_CD = ?`, schemeCode).
		Row()

	logger.Log(ctx).Debug("Query: ", zap.Any("Query: ", query))

	err := query.Scan(&transCount)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return false, err
	}

	var isTransactionExists = transCount > 0
	logger.Log(ctx).Debug("Result: ", zap.Any("Result: ", isTransactionExists))
	return isTransactionExists, nil
}

func (p portfolioDB) GetDPDetails(ctx context.Context, compCode int, schemeCode, folioNo, matchAccount string) (*models.DpDetails, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	var dpDetails *models.DpDetails

	query := p.oracle.WithContext(ctx).
		Select(`MF_UNB_DP_ID  AS "DpID"`,
			`MF_UNB_DP_ACC  AS "DpAccountNo"`,
			`NVL(MF_UNB_REINV_FLG, 'N') AS "ReinvestFlag"`).
		Table(`MF_UNIT_BAL`).
		Where(`MF_UNB_MATCH_ACC = ?`, matchAccount).
		Where(`MF_UNB_COMP_CD = ?`, compCode).
		Where(`MF_UNB_SCH_CD = ?`, schemeCode).
		Where(`MF_UNB_FOLIO = ?`, folioNo).
		Where(`ROWNUM = 1`)

	logger.Log(ctx).Debug("Query: ", zap.Any("Query: ", query))

	result := query.Find(&dpDetails)

	if result.Error != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		dpDetails.ReinvestFlag = "X"
	}

	logger.Log(ctx).Debug("Result: ", zap.Any("Result: ", dpDetails))
	return dpDetails, nil
}

func (p portfolioDB) GetDematUnblockedUnits(ctx context.Context, matchAccount, dpID, dpClientID, stockCode, isinNumber string) (*models.UnblockedUnits, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	var unblockedUnits *models.UnblockedUnits

	query := p.oracle.WithContext(ctx).
		Select(`NVL(MDBD_TOT_QTY,0) - NVL(MDBD_QTY_BLCKD,0)  AS "NoOfUnits"`).
		Table(`MDBD_DP_BLCK_DTLS`).
		Where(`MDBD_CLM_MTCH_ACCNT = ?`, matchAccount).
		Where(`MDBD_DP_ID = ?`, dpID).
		Where(`MDBD_DP_CLNT_ID = ?`, dpClientID).
		Where(`MDBD_STCK_CD = ?`, stockCode).
		Where(`MDBD_ISIN_NMBR = ?`, isinNumber).
		Where(`MDBD_TOT_QTY > 0`)

	logger.Log(ctx).Debug("Query: ", zap.Any("Query: ", query))

	result := query.Find(&unblockedUnits)

	if result.Error != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		unblockedUnits.ReinvestFlag = "X"
	}

	logger.Log(ctx).Debug("Result: ", zap.Any("unblockedUnits: ", unblockedUnits))
	return unblockedUnits, nil
}

func (p portfolioDB) GetUnblockedUnits(ctx context.Context, compCode int, schemeCode, folioNo, matchAccount string) (*models.UnblockedUnits, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	var unblockedUnits *models.UnblockedUnits

	query := p.oracle.WithContext(ctx).
		Select(`NVL(MF_UNB_NO_OF_UNITS, 0) AS "NoOfUnits"`,
			`TO_CHAR(MF_UNB_FOL_CREATION_DATE, 'dd-mm-yyyy') AS "FolioCreationDate"`,
			`NVL(MF_UNB_REINV_FLG, 'X') AS "ReinvestFlag"`).
		Table(`MF_UNIT_BAL`).
		Where(`MF_UNB_MATCH_ACC = ?`, matchAccount).
		Where(`MF_UNB_COMP_CD = ?`, compCode).
		Where(`MF_UNB_SCH_CD = ?`, schemeCode).
		Where(`MF_UNB_FOLIO = ?`, folioNo)

	logger.Log(ctx).Debug("Query: ", zap.Any("Query: ", query))

	result := query.Find(&unblockedUnits)

	if result.Error != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		unblockedUnits.ReinvestFlag = "X"
	}

	logger.Log(ctx).Debug("Result: ", zap.Any("unblockedUnits: ", unblockedUnits))
	return unblockedUnits, nil
}
