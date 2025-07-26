package db

import (
	"context"
	"kisaanSathi/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type allocatedFundsDB struct {
	oracle *gorm.DB
}

type AllocatedFundsStore interface {
	GetLinkedLimits(ctx context.Context, matchAccount string) (float64, error)
	GetTPALimits(ctx context.Context, matchAccount string) (float64, error)
	GetDepositLinkedLimits(ctx context.Context, matchAccount string) (float64, error)
}

func NewAllocatedFundsStore(oracle *gorm.DB) AllocatedFundsStore {
	return &allocatedFundsDB{oracle: oracle}
}

func (a allocatedFundsDB) GetLinkedLimits(ctx context.Context, matchAccount string) (float64, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	allocatedAmount := 0.0
	netBalance := 0.0
	debitCreditFlag := "D"
	query := a.oracle.WithContext(ctx).
		Select(`NVL(CLM_MF_ALLCTD_AMT, 0) AS "allocatedAmount"`,
			`NVL(MF_NT_BLNCS, 0) AS "netBalance"`,
			`NVL(MF_DB_CR_FLG, 'D') AS "debitCreditFlag"`).
		Table(`CLM_CLNT_MSTR, MF_TOT_BLNCS`).
		Where(`CLM_MTCH_ACCNT = MF_CLM_MTCH_ACCNT(+)`).
		Where(`CLM_MTCH_ACCNT = ?`, matchAccount).
		Row()
	err := query.Scan(&allocatedAmount, &netBalance, &debitCreditFlag)
	if err != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return 0, err
	}

	logger.Log(ctx).Debug("GetLinkedLimits: ", zap.Any("allocatedAmount: ", allocatedAmount), zap.Any("netBalance: ", netBalance), zap.Any("debitCreditFlag: ", debitCreditFlag))

	linkedLimits := 0.0
	if debitCreditFlag == "C" {
		linkedLimits = allocatedAmount
	} else {
		linkedLimits = allocatedAmount - netBalance
	}

	return linkedLimits, nil
}

func (a allocatedFundsDB) GetTPALimits(ctx context.Context, matchAccount string) (float64, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	allocatedAmount := 0.0
	netBalance := 0.0
	debitCreditFlag := "D"
	query := a.oracle.WithContext(ctx).
		Select(`NVL(CLM_MF_TPA_ALLCTD_AMT, 0) AS "allocatedAmount"`,
			`NVL(MTB_NT_BLNCS, 0) AS "netBalance"`,
			`NVL(MTB_DB_CR_FLG, 'D') AS "debitCreditFlag"`).
		Table(`CLM_CLNT_MSTR, MTB_MF_TPA_BLNCS`).
		Where(`CLM_MTCH_ACCNT = MTB_CLM_MTCH_ACCNT(+)`).
		Where(`CLM_MTCH_ACCNT = ?`, matchAccount).
		Row()
	err := query.Scan(&allocatedAmount, &netBalance, &debitCreditFlag)
	if err != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return 0, err
	}

	logger.Log(ctx).Debug("GetLinkedLimits: ", zap.Any("allocatedAmount: ", allocatedAmount), zap.Any("netBalance: ", netBalance), zap.Any("debitCreditFlag: ", debitCreditFlag))

	tpaLimits := 0.0
	if debitCreditFlag == "C" {
		tpaLimits = allocatedAmount
	} else {
		tpaLimits = allocatedAmount - netBalance
	}

	return tpaLimits, nil
}

func (a allocatedFundsDB) GetDepositLinkedLimits(ctx context.Context, matchAccount string) (float64, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	allocatedAmount := 0.0
	netBalance := 0.0
	debitCreditFlag := "D"
	query := a.oracle.WithContext(ctx).
		Select(`NVL(CLM_MF_ALLCTD_AMT_DEP, 0) AS "allocatedAmount"`,
			`NVL(MF_NT_BLNCS_DEP, 0) AS "netBalance"`,
			`NVL(MF_DB_CR_FLG_DEP, 'D') AS "debitCreditFlag"`).
		Table(`CLM_CLNT_MSTR, MF_TOT_BLNCS`).
		Where(`CLM_MTCH_ACCNT = MF_CLM_MTCH_ACCNT(+)`).
		Where(`CLM_MTCH_ACCNT = ?`, matchAccount).
		Row()
	err := query.Scan(&allocatedAmount, &netBalance, &debitCreditFlag)
	if err != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return 0, err
	}

	logger.Log(ctx).Debug("GetLinkedLimits: ", zap.Any("allocatedAmount: ", allocatedAmount), zap.Any("netBalance: ", netBalance), zap.Any("debitCreditFlag: ", debitCreditFlag))

	depositLinkedLimits := 0.0
	if debitCreditFlag == "C" {
		depositLinkedLimits = allocatedAmount
	} else {
		depositLinkedLimits = allocatedAmount - netBalance
	}

	return depositLinkedLimits, nil
}
