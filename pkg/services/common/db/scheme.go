package db

import (
	"context"
	"database/sql"
	"errors"
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/common/models"
	"kisaanSathi/pkg/services/common/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type schemeDB struct {
	oracle *gorm.DB
}

type SchemeFlagValidation interface {
	commonSchemeValidation(ctx context.Context, schemeFlags *models.SchemeFlags, isOfflineRequest, isSpecialInterval bool) error
	ValidateSchemeFlagsForPurchase(ctx context.Context, compCode int, schemeCode string, isOfflineRequest bool) (*models.SchemeFlags, error)
	ValidateSchemeFlagsForSIP(ctx context.Context, compCode int, schemeCode string, isOfflineRequest bool) (*models.SchemeFlags, error)
	ValidateSchemeFlagsForRedeem(ctx context.Context, compCode int, schemeCode string, isSpecialInterval bool) (*models.SchemeFlags, error)
	ValidateSchemeFlagsForSWP(ctx context.Context, compCode int, schemeCode string, isSpecialInterval bool) (*models.SchemeFlags, error)
	ValidateSchemeFlagsForSwitch(ctx context.Context, compCode int, schemeCode string, isSpecialInterval bool) (*models.SchemeFlags, error)
	ValidateSchemeFlagsForSTP(ctx context.Context, compCode int, schemeCode string, isSpecialInterval, isBoosterSTP, isDematHolding bool) (*models.SchemeFlags, error)
}

type SchemeStore interface {
	SchemeFlagValidation
	GetCompanyRegistrar(ctx context.Context, compCode int) (string, error)
	GetCompanyFolio(ctx context.Context, compCode int) (string, error)
	GetSchemeFlags(ctx context.Context, compCode int, schemeCode string) (*models.SchemeFlags, error)
	GetSchemeDetails(ctx context.Context, compCode int, schemeCode string) (*models.SchemeDetails, error)
	GetRedeemSchemeDetails(ctx context.Context, compCode int, schemeCode, holdingMode, accountType string) (*models.RedeemSchemeDetails, error)
	GetSchemeNavDetails(ctx context.Context, compCode int, schemeCode string) (*models.NavDetails, error)
	IsParamEATMEnabled(ctx context.Context) (bool, error)
	IsSchemeEATMEnabled(ctx context.Context, compCode int, schemeCode string) (bool, error)
}

func NewSchemeStore(oracle *gorm.DB) SchemeStore {
	return &schemeDB{oracle: oracle}
}

func (s schemeDB) GetSchemeFlags(ctx context.Context, compCode int, schemeCode string) (*models.SchemeFlags, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	var schemeFlags *models.SchemeFlags

	query := s.oracle.WithContext(ctx).
		Select(
			`NVL(MF_SCH_TM_MKT_STP_OUT_FLG,'N') AS "BoosterSTPFlag"`,
			`DECODE(NVL(MF_SCH_CLOSE_FLG, 'N'), 'Y', 'Close Ended', 'N', 'Open Ended') AS "CloseFlag"`,
			`NVL(MF_SCH_DIRECT_SCH_FLG, 'N') AS "DirectSchemeFlag"`,
			`NVL(MF_SCH_DR_FLG, 'N') AS "DRFlag"`,
			`NVL(MF_SCH_DIV_REINV_FLG, ' ') AS "DivReinvestFlag"`,
			`NVL(MF_SCH_ETF_FLG, 'N') AS "ETFFlag"`,
			`NVL(MF_SCH_FREED_SRC_FLG, 'N') AS "FreedomFlag"`,
			`NVL(MF_SCH_FREE_INSURE_FLG, 'N') AS "FreeInsureFlag"`,
			`NVL(MF_SCH_MUL_TRN_ALWD, 'N') AS MultiTransAllowedFlag`,
			`NVL(MF_SCH_OFLN_FLG, 'N') AS "OfflineFlag"`,
			`NVL(MF_SCH_ONLN_FLG, 'N') AS "OnlineFlag"`,
			`CASE
				WHEN SYSDATE > TO_DATE(TO_CHAR(MF_SCH_END_DATE, 'DD-Mon-YYYY')|| ' ' || MF_SCH_DISP_PUR_CUTOFF, 'DD-Mon-YYYY HH24:MI') THEN 'N'
				ELSE 'Y'
			END AS "PurchaseAllowedFlag"`,
			`NVL(MF_SCH_PURCHASE_FLG, 'N') AS "PurchaseFlag"`,
			`NVL(MF_REC_FLAG, ' ') AS "RecommendFlag"`,
			`CASE
				WHEN SYSDATE > TO_DATE(TO_CHAR(MF_SCH_END_DATE, 'DD-Mon-YYYY')|| ' ' || MF_SCH_DISP_REDEM_CUTOFF, 'DD-Mon-YYYY HH24:MI') THEN 'N'
				ELSE 'Y'
			END AS "RedeemAllowedFlag"`,
			`NVL(MF_SCH_REDEM_FLG, 'N') AS "RedeemFlag"`,
			`NVL(MF_SCH_RENEWAL_FLG, 'N') AS "RenewalFlag"`,
			`NVL(MF_SCH_AIP_FLG, 'N') AS "SIPFlag"`,
			`NVL(MF_SCH_SPCL_INTRVL_IND, 'N') AS "SpecialIntervalFlag"`,
			`NVL(MF_SCH_STEP_UP_FLG, 'N') AS "StepUpFlag"`,
			`NVL(MF_SCH_STP_OUT_FLG, 'N') AS "STPOutFlag"`,
			`NVL(MF_SCH_SWITCH_FLG, 'N') AS "SwitchFlag"`,
			`CASE
				WHEN SYSDATE > TO_DATE(TO_CHAR(MF_SCH_END_DATE, 'DD-Mon-YYYY')|| ' ' || MF_SCH_DISP_SISO_CUTOFF, 'DD-Mon-YYYY HH24:MI') THEN 'N'
				ELSE 'Y'
			END AS "SwitchAllowedFlag"`,
			`NVL(MF_SCH_AWP_FLG, 'N') AS "SWPFlag"`,
			`NVL(MF_SCH_TRGT_FUND_FLG, 'N') AS "TargetFundFlag"`).
		Table(`MF_SCHEME_MASTER, MF_RECOMMENDATIONS`).
		Where(`MF_SCH_COMP_CD = MF_COMPANY_CODE(+)`).
		Where(`MF_SCH_CD = MF_SCHEME_CODE(+)`).
		Where(`MF_SCH_COMP_CD = ?`, compCode).
		Where(`MF_SCH_CD = ?`, schemeCode)

	logger.Log(ctx).Debug("Query: ", zap.Any("Query: ", query))

	result := query.Find(&schemeFlags)

	if result.Error != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", utils.SchemeErrors.NoSchemeFound))
		return nil, errors.New(utils.SchemeErrors.NoSchemeFound)
	}

	logger.Log(ctx).Debug("Result: ", zap.Any("Result: ", schemeFlags))
	return schemeFlags, nil
}

func (s schemeDB) GetSchemeDetails(ctx context.Context, compCode int, schemeCode string) (*models.SchemeDetails, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	var schemeDetails *models.SchemeDetails

	query := s.oracle.WithContext(ctx).
		Select(
			`NVL(MF_SCH_MIN_PURCHASE_AMT, 0) AS "MinPurchaseAmount"`,
			`NVL(MF_SCH_PURCHASE_MULTI_AMT, 0) AS "MultiPurchaseAmount"`,
			`NVL(MF_SCH_MIN_INIT_PUR_AMT, 0) AS "MinSIPAmount"`,
			`NVL(MF_SCH_INIT_MULTI_AMT, 0) AS "MultiSIPAmount"`,
			`NVL(MF_SCH_MAX_SUB_AMT, 0) AS "MaxSubAmount"`,
			`NVL(TO_CHAR(MF_SCH_NFO_EXEC_DATE, 'dd/mm/yyyy'), 'Nil') AS "NFOExecDate"`,
			`NVL(MF_SCH_MAX_AIP_AMT, 0) AS MaxSipAmount`,
			`NVL(MF_SCH_DESC, ' ') AS "SchemeDesc"`,
			`UPPER(NVL(MF_SCH_TYPE, 'N')) AS "SchemeType"`).
		Table(`MF_SCHEME_MASTER`).
		Where(`MF_SCH_COMP_CD = ?`, compCode).
		Where(`MF_SCH_CD = ?`, schemeCode)

	logger.Log(ctx).Debug("Query: ", zap.Any("Query: ", query))

	result := query.Find(&schemeDetails)

	if result.Error != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", utils.SchemeErrors.NoSchemeFound))
		return nil, errors.New(utils.SchemeErrors.NoSchemeFound)
	}

	logger.Log(ctx).Debug("Result: ", zap.Any("Result: ", schemeDetails))
	return schemeDetails, nil
}

func (s schemeDB) GetRedeemSchemeDetails(ctx context.Context, compCode int, schemeCode, holdingMode, accountType string) (*models.RedeemSchemeDetails, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	var schemeDetails *models.RedeemSchemeDetails

	query := s.oracle.WithContext(ctx).
		Select(
			`NVL(MF_NAV_AMFI_CD,' ') AS "AmfiCode"`,
			`MF_SCH_DESC AS "SchemeDesc"`,
			`NVL(TO_CHAR(MF_NAV_DATE, 'dd-mm-yyyy'), 'Nil') AS "NavDate"`,
			`NVL(MF_NAV_NAV, 0) AS "NavValue"`,
			`NVL(MF_SCH_DISP_REDEM_CUTOFF, 0) AS "RedeemCutoffTime"`,
			`NVL(MF_SCH_DISP_SISO_CUTOFF, 0) AS "SisoCutoffTime"`,
			`NVL(MF_SCH_MIN_REDEM_AMT, 0) AS "MinRedeemAmount"`,
			`NVL(MF_SCH_RED_MULTI_AMT, 0) AS "MultiRedeemAmount"`,
			`NVL(MF_SCH_MIN_RED_UNITS, 0) AS "MinRedeemUnits"`,
			`NVL(MF_SCH_MULTI_RED_UNITS, 0) AS "MultiRedeemUnits"`,
			`NVL(MF_SCH_REMARKS, '0') AS "SchemeRemarks"`,
			`DECODE('`+holdingMode+`', 'D', ROUND(DECODE(NVL(MF_NAV_NAV, 0), 0, 0, NVL(MF_SCH_STP_OUT_MIN_AMT, 0)/ MF_NAV_NAV), 4), NVL(MF_SCH_STP_OUT_MIN_AMT, 0)) AS "StpMinAmount"`,
			`DECODE('`+holdingMode+`', 'D', ROUND(DECODE(NVL(MF_NAV_NAV, 0), 0, 0, NVL(MF_SCH_STP_OUT_MAX_AMT, 0)/ MF_NAV_NAV), 4), NVL(MF_SCH_STP_OUT_MAX_AMT, 0)) AS "StpMaxAmount"`,
			`NVL(MF_SCH_STP_MIN_HLDNG_AMT, 0) AS "StpMinHoldingAmount"`,
			`NVL(DECODE('`+accountType+`', 'A', MF_SCH_NRE, 'B', MF_SCH_NPNRE, 'C', MF_SCH_NRO, 'D', MF_SCH_NPNRO, 'O', MF_SCH_RI), 'Y') AS "AccountTypeAllowedFlag"`).
		Table(`MF_NAVS, MF_SCHEME_MASTER`).
		Where(`MF_NAV_COMP_CD (+) = MF_SCH_COMP_CD`).
		Where(`MF_NAV_SCH_CD (+) = MF_SCH_CD`).
		Where(`MF_SCH_START_DATE <= SYSDATE`).
		Where(`MF_SCH_END_DATE + 1 >= SYSDATE`).
		Where(`MF_SCH_COMP_CD = ?`, compCode).
		Where(`MF_SCH_CD = ?`, schemeCode).
		Order(`MF_SCH_DESC`)

	logger.Log(ctx).Debug("Query: ", zap.Any("Query: ", query))

	result := query.Find(&schemeDetails)

	if result.Error != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Log(ctx).Error(utils.SchemeErrors.NoSchemeFound)
		return nil, errors.New(utils.SchemeErrors.NoSchemeFound)
	}

	logger.Log(ctx).Debug("Result: ", zap.Any("Result: ", schemeDetails))
	return schemeDetails, nil
}

func (s schemeDB) GetSchemeNavDetails(ctx context.Context, compCode int, schemeCode string) (*models.NavDetails, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	var navDetails *models.NavDetails

	query := s.oracle.WithContext(ctx).
		Select(`NVL(MF_NAV_AMFI_CD,' ') AS "AmfiCode"`,
			`NVL(TO_CHAR(MF_NAV_DATE,'dd-mm-yyyy'), 'Nil') AS "NavDate"`,
			`NVL(MF_NAV_NAV, 0) AS "NavValue"`).
		Table(`MF_NAVS`).
		Where(`MF_NAV_COMP_CD = ?`, compCode).
		Where(`MF_NAV_SCH_CD = ?`, schemeCode)

	logger.Log(ctx).Debug("Query: ", zap.Any("Query: ", query))

	result := query.Find(&navDetails)

	if result.Error != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Log(ctx).Error(utils.SchemeErrors.NoNavFound)
		return nil, errors.New(utils.SchemeErrors.NoNavFound)
	}

	logger.Log(ctx).Debug("Result: ", zap.Any("Result: ", navDetails))
	return navDetails, nil
}

func (s schemeDB) GetCompanyRegistrar(ctx context.Context, compCode int) (string, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	registrar := ""
	query := s.oracle.WithContext(ctx).
		Select(`UPPER(NVL(MF_COMP_REGISTRAR,'-')) AS "registrar"`).
		Table(`MF_COMPANIES`).
		Where(`MF_COMP_CD = ?`, compCode).
		Row()
	err := query.Scan(&registrar)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return "", err
	}

	if errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error(utils.SchemeErrors.NoRegistrarFound)
		return "", errors.New(utils.SchemeErrors.NoRegistrarFound)
	}

	logger.Log(ctx).Debug("registrar: ", zap.Any("Result: ", registrar))

	return registrar, nil
}

func (s schemeDB) GetCompanyFolio(ctx context.Context, compCode int) (string, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	compFolioFlag := "N"
	query := s.oracle.WithContext(ctx).
		Select(`NVL(MF_COMP_FOL_FLG, 'N') AS "compFolioFlag"`).
		Table(`MF_COMPANIES`).
		Where(`MF_COMP_CD = ?`, compCode).
		Row()
	err := query.Scan(&compFolioFlag)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return "", err
	}

	if errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error(utils.SchemeErrors.NoCompanyFound)
		return "", errors.New(utils.SchemeErrors.NoCompanyFound)
	}

	logger.Log(ctx).Debug("Result: ", zap.Any("compFolioFlag: ", compFolioFlag))

	return compFolioFlag, nil
}

func (s schemeDB) IsParamEATMEnabled(ctx context.Context) (bool, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	isParamEATMEnabled := false

	eATMFlag := "N"
	query := s.oracle.WithContext(ctx).
		Select(`NVL(MF_PAR_EATM_FLG, 'N') AS "eATMFlag"`).
		Table(`MF_PARAM`).
		Row()

	err := query.Scan(&eATMFlag)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return false, err
	}

	isParamEATMEnabled = eATMFlag == "Y"

	logger.Log(ctx).Debug("Result: ", zap.Any("isEATMEnabled: ", isParamEATMEnabled))

	return isParamEATMEnabled, nil
}

func (s schemeDB) IsSchemeEATMEnabled(ctx context.Context, compCode int, schemeCode string) (bool, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	isSchemeEATMEnabled := false

	compEATMFlag := "N"
	schemeEATMFlag := "N"
	categoryEATMFlag := "N"
	query := s.oracle.WithContext(ctx).
		Select(
			`NVL(MF_COMP_EATM_FLG, 'N') AS "compEATMFlag"`,
			`NVL(MF_SCH_EATM_FLG, 'N') AS "schemeEATMFlag"`,
			`NVL(MF_CAT_EATM_FLG, 'N') AS "categoryEATMFlag"`).
		Table(`MF_COMPANIES, MF_SCHEME_MASTER, MF_CATEGORY_MASTER`).
		Where(`MF_COMP_CD = MF_SCH_COMP_CD`).
		Where(`MF_SCH_CAT_CD = MF_CAT_CD`).
		Where(`MF_COMP_CD = ?`, compCode).
		Where(`MF_SCH_CD = ?`, schemeCode).
		Row()

	err := query.Scan(&compEATMFlag, &schemeEATMFlag, &categoryEATMFlag)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return false, err
	}

	isSchemeEATMEnabled = compEATMFlag == "Y" && schemeEATMFlag == "Y" && categoryEATMFlag == "Y"

	logger.Log(ctx).Debug("Result: ", zap.Any("isSchemeEATMEnabled: ", isSchemeEATMEnabled))

	return isSchemeEATMEnabled, nil
}
