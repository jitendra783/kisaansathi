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

type userDB struct {
	oracle *gorm.DB
}

type UserStore interface {
	GetUserDetails(ctx context.Context, userID, matchAccount, isAgent string) (*models.UserInfo, error)
	getPersonalInfo(ctx context.Context, matchAccount string) (*models.PersonalInfo, error)
	getExtraUserInfo(ctx context.Context, matchAccount string) (*models.ExtraInfo, error)
	GetPrivacyInfo(ctx context.Context, userID, panNo string, isModifyRequired bool) (*models.PrivacyInfo, error)
	GetCustodianFlag(ctx context.Context, matchAccount string) (string, error)
	GetDematAccount(ctx context.Context, matchAccount string) (string, error)
	GetSipHealthFlag(ctx context.Context, matchAccount string) (string, error)
	GetEBAUploadDate(ctx context.Context, userID string) (string, error)
	IsD2UEnabled(ctx context.Context, matchAccount string) (bool, error)
	IsD2UActive(ctx context.Context, matchAccount string) (bool, error)
	euinNo
	IsEATMEnabled(ctx context.Context, matchAccount string) (bool, error)
	GetAccountType(ctx context.Context, matchAccount string) (string, error)
}

type euinNo interface {
	GetEuinNo(ctx context.Context, userID, matchAccount, callAndTradeUserID string, isD2uUser bool) (string, error)
	getEuinNoForBPID(ctx context.Context, bpID string) (string, error)
	getEuinNoForBusinessPartner(ctx context.Context, bpID, matchAccount string) (string, error)
	getEuinNoForRegularUser(ctx context.Context, matchAccount string, isD2uUser bool) (string, error)
	getEuinNoForCallAndTradeUser(ctx context.Context, callAndTradeUserID string) (string, error)
}

func (u *userDB) GetSipHealthFlag(ctx context.Context, matchAccount string) (string, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	sipHealthFlag := ""
	query := u.oracle.WithContext(ctx).
		Select(`NVL(CSD_MF_SNH_FLG,'N') AS "sipHealthFlag"`).
		Table(`CSD_CLNT_SPT_DTLS`).
		Where(`CSD_CLM_MTCH_ACCNT = ?`, matchAccount).
		Row()
	err := query.Scan(&sipHealthFlag)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return "", err
	}

	if errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("no flag found")
		return "N", nil
	}

	logger.Log(ctx).Debug("Result: ", zap.Any("sipHealthFlag: ", sipHealthFlag))

	return sipHealthFlag, nil
}

func (u *userDB) GetDematAccount(ctx context.Context, matchAccount string) (string, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	dematAccountNo := ""
	query := u.oracle.WithContext(ctx).
		Select(`CLM_MF_DMAT_ACC AS "dematAccountNo"`).
		Table(`CLM_CLNT_MSTR`).
		Where(`CLM_MTCH_ACCNT = ?`, matchAccount).
		Row()
	err := query.Scan(&dematAccountNo)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return "", err
	}

	if errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("no demat account")
		return "", errors.New("no demat account")
	}

	logger.Log(ctx).Debug("dematAccountNo: ", zap.Any("Result: ", dematAccountNo))

	return dematAccountNo, nil
}

func (u *userDB) GetCustodianFlag(ctx context.Context, matchAccount string) (string, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	custodianFlag := "N"
	query := u.oracle.WithContext(ctx).
		Select(`NVL(ICD_MFCUSTODIAN_FLG, 'N') AS "custodianFlag"`).
		Table(`ICD_INFO_CLIENT_DTLS, IAI_INFO_ACCOUNT_INFO`).
		Where(`ICD_SERIAL_NO = IAI_SERIAL_NO`).
		Where(`IAI_MATCH_ACCOUNT_NO = ?`, matchAccount).
		Row()
	err := query.Scan(&custodianFlag)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return "", err
	}

	logger.Log(ctx).Debug("custodianFlag: ", zap.Any("Result: ", custodianFlag))

	return custodianFlag, nil
}

func (u *userDB) GetUserDetails(ctx context.Context, userID, matchAccount, isAgent string) (*models.UserInfo, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	isBackOfficeUser := utils.Strings(userID).IsBackOfficeUser()
	isAgentUser := utils.Strings(isAgent).IsAgent()

	var userInfo models.UserInfo

	if !isBackOfficeUser {
		personalInfo, err := u.getPersonalInfo(ctx, matchAccount)
		if err != nil {
			return nil, err
		}
		userInfo.PersonalInfo = *personalInfo
	}

	extraInfo, err := u.getExtraUserInfo(ctx, matchAccount)
	if err != nil {
		return nil, err
	}
	userInfo.CustomerInfo.IsRICustomer = extraInfo.CustomerInfo.IsRICustomer

	if isAgentUser {
		userInfo.CustomerInfo.PanNo = extraInfo.CustomerInfo.PanNo

		privacyInfo, err := u.GetPrivacyInfo(ctx, userID, extraInfo.PanNo, false)
		if err != nil {
			return nil, err
		}

		if privacyInfo.PrivacyOpted {
			userInfo.ExtraPersonalInfo = extraInfo.ExtraPersonalInfo
		}
	}

	logger.Log(ctx).Debug("userInfo - ", zap.Any("userInfo: ", userInfo))
	return &userInfo, nil
}

func (u *userDB) getPersonalInfo(ctx context.Context, matchAccount string) (*models.PersonalInfo, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	var response *models.PersonalInfo

	query := u.oracle.WithContext(ctx).
		Select(
			`NVL(TRIM(IPD_FIRST_NAME), ' ') || ' ' || NVL(TRIM(IPD_MIDDLE_NAME), ' ') || ' ' || NVL(TRIM(IPD_LAST_NAME), ' ') AS "FullName"`,
			`NVL(IPD_EMAIL, 'notprovided@notprovided.com') AS "Email"`,
			`NVL(IAD_MOBILE, '*') AS "MobileNo"`,
			`NVL(IPD_SEX, 'O') AS "Gender"`,
		).
		Table(`IPD_INFO_PERSONAL_DTLS, IAI_INFO_ACCOUNT_INFO, IAD_INFO_ADDRESS_DTLS`).
		Where(`IAI_SERIAL_NO = IPD_SERIAL_NO`).
		Where(`IAD_SERIAL_NO = IPD_SERIAL_NO`).
		Where(`IPD_TYPE = 'APPLICANT'`).
		Where(`IAD_ADDRESS_TYPE = 'APPLICANT_CORR'`).
		Where(`IAI_MATCH_ACCOUNT_NO = ?`, matchAccount)

	logger.Log(ctx).Debug("Query: ", zap.Any("Query: ", query))

	result := query.Find(&response)

	if result.Error != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Log(ctx).Error("No results found ")
		return nil, errors.New("no user found")
	}

	logger.Log(ctx).Debug("Result: ", zap.Any("Result: ", response))
	return response, nil
}

func (u *userDB) getExtraUserInfo(ctx context.Context, matchAccount string) (*models.ExtraInfo, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	var response *models.ExtraInfo

	query := u.oracle.WithContext(ctx).
		Select(
			`NVL(IID_NO, '*') AS "PanNo"`,
			`DECODE(ICD_CUST_TYPE, 'NRI','N', 'COR', 'N','HUF','N', 'PAR', 'N', 'Y') AS "IsRICustomer"`,
			`NVL(IPD_FIRST_NAME, '*') AS "FirstName"`,
			`NVL(IPD_MIDDLE_NAME, '*') AS "MiddleName"`,
			`NVL(IPD_LAST_NAME, '*') AS "LastName"`,
			`NVL(TO_CHAR(IPD_DOB, 'dd-Mon-yyyy'), '*') AS "DateOfBirth"`,
			`NVL(IAD_TEL_RES, '*') AS "ResidentialNo"`,
			`NVL(IAD_TEL_OFF, '*') AS "OfficialNo"`,
			`NVL(IPD_EMAIL, '*') AS "EmailAddress"`,
		).
		Table(`ICD_INFO_CLIENT_DTLS, IPD_INFO_PERSONAL_DTLS, IAD_INFO_ADDRESS_DTLS, IID_INFO_IDENTIFICATION_DTLS, UAC_USR_ACCNTS`).
		Where(`ICD_SERIAL_NO = IPD_SERIAL_NO`).
		Where(`ICD_SERIAL_NO = IAD_SERIAL_NO`).
		Where(`ICD_SERIAL_NO = IID_SERIAL_NO`).
		Where(`ICD_USER_ID = UAC_USR_ID`).
		Where(`IPD_TYPE = 'APPLICANT'`).
		Where(`IAD_ADDRESS_TYPE = 'APPLICANT_CORR'`).
		Where(`IID_TYPE = 'PAN'`).
		Where(`UAC_CLM_MTCH_ACCNT = ?`, matchAccount)

	logger.Log(ctx).Debug("Query: ", zap.Any("Query: ", query))

	result := query.Find(&response)

	if result.Error != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", result.Error))
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Log(ctx).Error("No results found ")
		return nil, errors.New("no user found")
	}

	if response.PanNo == "*" {
		logger.Log(ctx).Error("pan not found")
		return nil, errors.New("pan not found")
	}

	if response.DateOfBirth == "*" {
		logger.Log(ctx).Error("date of birth not found")
		return nil, errors.New("date of birth not found")
	}

	logger.Log(ctx).Debug("Result: ", zap.Any("Result: ", response))
	return response, nil
}

func (u *userDB) GetPrivacyInfo(ctx context.Context, userID, panNo string, isModifyRequired bool) (*models.PrivacyInfo, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	response := &models.PrivacyInfo{
		PrivacyOpted:  false,
		ModifyAllowed: "N",
	}

	var nodeID string
	query := u.oracle.WithContext(ctx).
		Select(`NVL(CMM_NODE_ID, '*') AS "nodeID"`).
		Table(`CMM_CLNT_MAP_MSTR`).
		Where(`CMM_OFF_MAP_FLG = 'Y'`).
		Where(`CMM_STTS = 'A'`).
		Where(`CMM_CLN_PAN_NO = ?`, panNo).
		Where(`CMM_ENTRY_DATE = (SELECT MIN(CMM_ENTRY_DATE) FROM CMM_CLNT_MAP_MSTR WHERE CMM_CLN_PAN_NO = ?)`, panNo).
		Row()
	err := query.Scan(&nodeID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: no entry found for this pan in CMM table")
		return response, errors.New("no entry found for this pan in CMM table")
	}

	logger.Log(ctx).Debug("Node Result: ", zap.Any("Result: ", nodeID))

	var transCount uint16
	query = u.oracle.WithContext(ctx).
		Select(`COUNT(*) AS "transCount"`).
		Table(`MF_TRANSACTIONS`).
		Where(`MF_TRN_STATUS_CD = 'E'`).
		Where(`MF_TRN_BP_ID = ?`, userID).
		Where(`MF_TRN_PAN_NO = ?`, panNo).
		Row()
	err = query.Scan(&transCount)

	if err != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return nil, err
	}

	logger.Log(ctx).Debug("Trans Count Result: ", zap.Any("Result: ", transCount))

	if userID == nodeID || transCount > 0 {
		response.PrivacyOpted = true

		if isModifyRequired {
			response.ModifyAllowed = "U"
		}
	} else {
		var nodeCount uint16
		query = u.oracle.WithContext(ctx).
			Select(`COUNT(*) AS "nodeCount"`).
			Table(`CMM_CLNT_MAP_MSTR`).
			Where(`CMM_OFF_MAP_FLG = 'Y'`).
			Where(`CMM_STTS = 'A'`).
			Where(`CMM_NODE_ID = ?`, userID).
			Where(`CMM_CLN_PAN_NO = ?`, panNo).
			Row()
		err = query.Scan(&nodeCount)

		if err != nil {
			logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
			return nil, err
		}

		if nodeCount <= 0 {
			response.ModifyAllowed = "M"
		}

		logger.Log(ctx).Debug("Node Count Result: ", zap.Any("Result: ", nodeCount))
	}

	return response, nil
}

func (u *userDB) GetEuinNo(ctx context.Context, userID, matchAccount, callAndTradeUserID string, isD2uUser bool) (string, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	isBusinessPartner := utils.Strings(userID).IsBusinessPartner()
	if isBusinessPartner {
		return u.getEuinNoForBusinessPartner(ctx, userID, matchAccount)
	} else if callAndTradeUserID != "" {
		return u.getEuinNoForCallAndTradeUser(ctx, callAndTradeUserID)
	} else {
		return u.getEuinNoForRegularUser(ctx, matchAccount, isD2uUser)
	}

}

func (u *userDB) getEuinNoForCallAndTradeUser(ctx context.Context, callAndTradeUserID string) (string, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	euinNo := ""

	query := u.oracle.WithContext(ctx).
		Select(`UPPER(EED_EUIN_NO) AS "euinNo"`).
		Table(`BPA_EMP_EUIN_DTLS, HEE_HR_EMP_EXTRCT`).
		Where(`HEE_EMPLID = EED_EMP_ID`).
		Where(`HEE_EMPL_STATUS = DECODE(HEE_SOURCE, 'ICICI SECURITIES', 'Y', 'ICICI BANK', 'A')`).
		Where(`EED_EUIN_EXP_DT > TRUNC(SYSDATE)`).
		Where(`NVL(HEE_RELATION_CD, 0) = 0`).
		Where(`(HEE_SOURCE = 'ICICI SECURITIES' OR HEE_SOURCE = 'ICICI BANK')`).
		Where(`UPPER(EED_EMP_ID) = UPPER(?)`, callAndTradeUserID).
		Row()
	err := query.Scan(&euinNo)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return "", err
	}

	return euinNo, nil
}

func (u *userDB) getEuinNoForRegularUser(ctx context.Context, matchAccount string, isD2uUser bool) (string, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	if isD2uUser {
		return "", nil
	}

	euinNo := ""

	var bpID string
	query := u.oracle.WithContext(ctx).
		Select(`NVL(CLM_BP_ID, 'N') AS "bpID"`).
		Table(`CLM_CLNT_MSTR`).
		Where(`CLM_MTCH_ACCNT = ?`, matchAccount).
		Row()
	err := query.Scan(&bpID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return "", err
	}

	if bpID != "" && bpID != "N" {
		euinNo, err = u.getEuinNoForBPID(ctx, bpID)
		if err != nil {
			logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
			return "", err
		}
	}

	return euinNo, nil
}

func (u *userDB) getEuinNoForBusinessPartner(ctx context.Context, bpID string, matchAccount string) (string, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	euinNo := ""

	var familyType string
	query := u.oracle.WithContext(ctx).
		Select(`NVL(ICAD_FAMILY_TYP, 'X') AS "familyType"`).
		Table(`ICAD_INFO_CLIENT_ADDL_DTLS, IAI_INFO_ACCOUNT_INFO`).
		Where(`ICAD_SERIAL_NO = IAI_SERIAL_NO`).
		Where(`IAI_MATCH_ACCOUNT_NO = ?`, matchAccount).
		Row()
	err := query.Scan(&familyType)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return "", err
	}

	if familyType != "FAN" {
		euinNo, err = u.getEuinNoForBPID(ctx, bpID)
		if err != nil {
			logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
			return "", err
		}
	}

	return euinNo, nil
}

func (u *userDB) getEuinNoForBPID(ctx context.Context, bpID string) (string, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	euinNo := ""

	query := u.oracle.WithContext(ctx).
		Select(`CTCL_UIDNO AS "euinNo"`).
		Table(`BPA_CTCL_DETAILS`).
		Where(`CTCL_SEGID = '7'`).
		Where(`CTCL_STATUS = 1`).
		Where(`CTCL_UID_EXPIRYDATE > TRUNC(SYSDATE)`).
		Where(`CTCL_NODECODE IN (SELECT TO_CHAR(ND_ID) FROM BPA_NODE_DETAILS WHERE ND_CODE = ?)`, bpID).
		Row()
	err := query.Scan(&euinNo)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return "", err
	}

	return euinNo, nil
}

func (u *userDB) GetEBAUploadDate(ctx context.Context, userID string) (string, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	ebaUploadDate := ""

	query := u.oracle.WithContext(ctx).
		Select(`TO_CHAR(TRUNC(ICD_EBA_UPLOAD_DT), 'dd/mm/yyyy') AS "ebaUploadDate"`).
		Table(`ICD_INFO_CLIENT_DTLS, ICAD_INFO_CLIENT_ADDL_DTLS`).
		Where(`ICAD_SERIAL_NO = ICD_SERIAL_NO`).
		Where(`ICD_USER_ID = ?`, userID).
		Row()

	err := query.Scan(&ebaUploadDate)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return "", err
	}

	logger.Log(ctx).Debug("Result: ", zap.Any("ebaUploadDate: ", ebaUploadDate))

	return ebaUploadDate, nil
}

func (u *userDB) IsD2UEnabled(ctx context.Context, matchAccount string) (bool, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	isD2uEnabled := false

	rowCount := 0
	query := u.oracle.WithContext(ctx).
		Select(`COUNT(*) AS "rowCount"`).
		Table(`DMM_D2U_MATCH_MPPNG_MSTR`).
		Where(`DMM_REG_END_DT is null`).
		Where(`DMM_MATCH_ACC = ?`, matchAccount).
		Row()

	err := query.Scan(&rowCount)

	if err != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return false, err
	}

	isD2uEnabled = rowCount > 0

	logger.Log(ctx).Debug("Result: ", zap.Any("isD2uEnabled: ", isD2uEnabled))

	return isD2uEnabled, nil
}

func (u *userDB) IsD2UActive(ctx context.Context, matchAccount string) (bool, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	isD2uActive := false

	rowCount := 0
	query := u.oracle.WithContext(ctx).
		Select(`COUNT(*) AS "rowCount"`).
		Table(`DCM_D2U_CLNT_MSTR, UAC_USR_ACCNTS`).
		Where(`DCM_USER_ID = UAC_USR_ID`).
		Where(`DCM_PLAN_ACTIVE_FLAG = 'A'`).
		Where(`UAC_CLM_MTCH_ACCNT = ?`, matchAccount).
		Where(`ROWNUM = 1`).
		Row()

	err := query.Scan(&rowCount)

	if err != nil {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return false, err
	}

	isD2uActive = rowCount > 0

	logger.Log(ctx).Debug("Result: ", zap.Any("isD2uActive: ", isD2uActive))

	return isD2uActive, nil
}

func (u *userDB) IsEATMEnabled(ctx context.Context, matchAccount string) (bool, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	isEATMEnabled := false

	eATMFlag := "N"
	query := u.oracle.WithContext(ctx).
		Select(`NVL(CSD_MF_EATM_FLG, 'N') AS "eATMFlag"`).
		Table(`CSD_CLNT_SPT_DTLS`).
		Where(`CSD_CLM_MTCH_ACCNT = ?`, matchAccount).
		Row()

	err := query.Scan(&eATMFlag)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return false, err
	}

	isEATMEnabled = eATMFlag == "Y"

	logger.Log(ctx).Debug("Result: ", zap.Any("isEATMEnabled: ", isEATMEnabled))

	return isEATMEnabled, nil
}

func (u *userDB) GetAccountType(ctx context.Context, matchAccount string) (string, error) {
	logger.Log(ctx).Debug("START")
	defer logger.Log(ctx).Debug("END")

	accountType := ""
	query := u.oracle.WithContext(ctx).
		Select(`DECODE(IAI_TYPE, 'NRE_PINS', 'A', 'NRE_NON_PINS', 'B', 'NRO_PINS', 'C', 'NRO_NON_PINS', 'D', 'O') AS "accountType"`).
		Table(`IAI_INFO_ACCOUNT_INFO`).
		Where(`IAI_MATCH_ACCOUNT_NO = ?`, matchAccount).
		Row()
	err := query.Scan(&accountType)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error("Error: ", zap.Any("err: ", err))
		return "", err
	}

	if errors.Is(err, sql.ErrNoRows) {
		logger.Log(ctx).Error(utils.SchemeErrors.NoUserFound)
		return "", errors.New(utils.SchemeErrors.NoUserFound)
	}

	return accountType, nil
}

func NewUserStore(oracle *gorm.DB) UserStore {
	return &userDB{oracle: oracle}
}
