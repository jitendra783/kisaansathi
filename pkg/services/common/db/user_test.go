package db

import (
	"context"
	"database/sql"
	"errors"
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/utils"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserSuite struct {
	suite.Suite
	ctx       context.Context
	sqlDB     *sql.DB
	gormDB    *gorm.DB
	sqlMock   sqlmock.Sqlmock
	userStore UserStore
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (suite *UserSuite) SetupSuite() {
	logger.LoggerInit("", -1)

	suite.ctx = context.TODO()
	suite.sqlDB, suite.gormDB, suite.sqlMock = utils.NewMockDB()
	suite.userStore = NewUserStore(suite.gormDB)
}

func (suite *UserSuite) TearDownSuite() {
}

func (suite *UserSuite) SetupTest() {
}

func (suite *UserSuite) TearDownTest() {
}

func (suite *UserSuite) TestGetDematAccount() {
	testCases := []struct {
		desc           string
		mockInput      *sqlmock.Rows
		expectedError  string
		expectedOutput string
	}{
		{
			desc:           "SQLError",
			mockInput:      nil,
			expectedError:  "ORA Error",
			expectedOutput: "",
		}, {
			desc:           "NoRowsError",
			mockInput:      sqlmock.NewRows([]string{"dematAccountNo"}),
			expectedError:  "no demat account",
			expectedOutput: "",
		}, {
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"dematAccountNo"}).AddRow("12334"),
			expectedError:  "",
			expectedOutput: "12334",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM CLM_CLNT_MSTR WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM CLM_CLNT_MSTR WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.userStore.GetDematAccount(suite.ctx, "12345678")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput, testCase.expectedOutput)
			}

		})
	}
}

func (suite *UserSuite) TestGetCustodianFlag_ReturnSQLError() {
	// Mocking and Setting Expected Result
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, IAI_INFO_ACCOUNT_INFO WHERE (.+)$").
		WillReturnError(errors.New("ORA Error"))

	// Triggering Function
	_, err := suite.userStore.GetCustodianFlag(suite.ctx, "000000000")

	// Validations
	suite.EqualError(err, "ORA Error")
}

func (suite *UserSuite) TestGetCustodianFlag_ReturnDefaultValue_WhenNoRecordsFound() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"custodianFlag"})
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, IAI_INFO_ACCOUNT_INFO WHERE (.+)$").
		WillReturnRows(expectedOutput)

	// Triggering Function
	custodianFlag, err := suite.userStore.GetCustodianFlag(suite.ctx, "000000000")

	// Validations
	suite.NoError(err)
	suite.Equal(custodianFlag, "N")
}

func (suite *UserSuite) TestGetCustodianFlag_ReturnFlagValue_WhenRecordsFound() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"custodianFlag"}).AddRow("Y")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, IAI_INFO_ACCOUNT_INFO WHERE (.+)$").
		WillReturnRows(expectedOutput)

	// Triggering Function
	custodianFlag, err := suite.userStore.GetCustodianFlag(suite.ctx, "000000000")

	// Validations
	suite.NoError(err)
	suite.Equal(custodianFlag, "Y")
}

func (suite *UserSuite) TestGetPersonalInfo_ReturnSQLError() {
	// Mocking and Setting Expected Result
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM IPD_INFO_PERSONAL_DTLS, IAI_INFO_ACCOUNT_INFO, IAD_INFO_ADDRESS_DTLS WHERE (.+)$").
		WillReturnError(errors.New("ORA Error"))

	// Triggering Function
	_, err := suite.userStore.getPersonalInfo(suite.ctx, "000000000")

	// Validations
	suite.EqualError(err, "ORA Error")
}

func (suite *UserSuite) TestGetPersonalInfo_ReturnNoUserFoundError() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"FullName", "Email", "MobileNo", "Gender"})
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM IPD_INFO_PERSONAL_DTLS, IAI_INFO_ACCOUNT_INFO, IAD_INFO_ADDRESS_DTLS WHERE (.+)$").
		WillReturnRows(expectedOutput)

	// Triggering Function
	_, err := suite.userStore.getPersonalInfo(suite.ctx, "000000000")

	// Validations
	suite.EqualError(err, "no user found")
}

func (suite *UserSuite) TestGetPersonalInfo_ReturnValidUser() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"FullName", "Email", "MobileNo", "Gender"}).AddRow("Test User", "testemail@mail.com", "1234567890", "M")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM IPD_INFO_PERSONAL_DTLS, IAI_INFO_ACCOUNT_INFO, IAD_INFO_ADDRESS_DTLS WHERE (.+)$").
		WillReturnRows(expectedOutput)

	// Triggering Function
	actualOutput, err := suite.userStore.getPersonalInfo(suite.ctx, "111111111")

	// Validations
	suite.NoError(err)
	suite.Equal("Test User", actualOutput.FullName)
	suite.Equal("testemail@mail.com", actualOutput.Email)
	suite.Equal("M", actualOutput.Gender)
	suite.Equal("1234567890", actualOutput.MobileNo)
}

func (suite *UserSuite) TestGetExtraUserInfo_ReturnSQLError() {
	// Mocking and Setting Expected Result
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, IPD_INFO_PERSONAL_DTLS, IAD_INFO_ADDRESS_DTLS, IID_INFO_IDENTIFICATION_DTLS, UAC_USR_ACCNTS WHERE (.+)$").
		WillReturnError(errors.New("ORA Error"))

	// Triggering Function
	_, err := suite.userStore.getExtraUserInfo(suite.ctx, "000000000")

	// Validations
	suite.EqualError(err, "ORA Error")
}

func (suite *UserSuite) TestGetExtraUserInfo_ReturnNoUserFoundError() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"PanNo", "DateOfBirth"})
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, IPD_INFO_PERSONAL_DTLS, IAD_INFO_ADDRESS_DTLS, IID_INFO_IDENTIFICATION_DTLS, UAC_USR_ACCNTS WHERE (.+)$").
		WillReturnRows(expectedOutput)

	// Triggering Function
	_, err := suite.userStore.getExtraUserInfo(suite.ctx, "000000000")

	// Validations
	suite.EqualError(err, "no user found")
}

func (suite *UserSuite) TestGetExtraUserInfo_ReturnPANError() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"PanNo"}).AddRow("*")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, IPD_INFO_PERSONAL_DTLS, IAD_INFO_ADDRESS_DTLS, IID_INFO_IDENTIFICATION_DTLS, UAC_USR_ACCNTS WHERE (.+)$").
		WillReturnRows(expectedOutput)

	// Triggering Function
	_, err := suite.userStore.getExtraUserInfo(suite.ctx, "000000000")

	// Validations
	suite.EqualError(err, "pan not found")
}

func (suite *UserSuite) TestGetExtraUserInfo_ReturnDOBError() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"DateOfBirth"}).AddRow("*")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, IPD_INFO_PERSONAL_DTLS, IAD_INFO_ADDRESS_DTLS, IID_INFO_IDENTIFICATION_DTLS, UAC_USR_ACCNTS WHERE (.+)$").
		WillReturnRows(expectedOutput)

	// Triggering Function
	_, err := suite.userStore.getExtraUserInfo(suite.ctx, "000000000")

	// Validations
	suite.EqualError(err, "date of birth not found")
}

func (suite *UserSuite) TestGetExtraUserInfo_ReturnValidExtraUserInfo() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"PanNo", "DateOfBirth", "FirstName", "IsRICustomer"}).AddRow("AAAAA0000A", "01-Jan-1990", "Test User", "Y")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, IPD_INFO_PERSONAL_DTLS, IAD_INFO_ADDRESS_DTLS, IID_INFO_IDENTIFICATION_DTLS, UAC_USR_ACCNTS WHERE (.+)$").
		WillReturnRows(expectedOutput)

	// Triggering Function
	actualOutput, err := suite.userStore.getExtraUserInfo(suite.ctx, "000000000")

	// Validations
	suite.NoError(err)
	suite.Equal("Test User", actualOutput.FirstName)
	suite.Equal("AAAAA0000A", actualOutput.PanNo)
	suite.Equal("01-Jan-1990", actualOutput.DateOfBirth)
	suite.Equal("Y", actualOutput.IsRICustomer)
}

func (suite *UserSuite) TestGetUserDetails_ReturnNoPersonalInfoFoundError() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"FullName", "Email", "MobileNo", "Gender"})
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM IPD_INFO_PERSONAL_DTLS, IAI_INFO_ACCOUNT_INFO, IAD_INFO_ADDRESS_DTLS WHERE (.+)$").
		WillReturnRows(expectedOutput)

	// Triggering Function
	_, err := suite.userStore.GetUserDetails(suite.ctx, "TestUser", "12345678", "N")

	// Validations
	suite.EqualError(err, "no user found")
}

func (suite *UserSuite) TestGetUserDetails_ReturnNoPANError() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"FullName", "Email", "MobileNo", "Gender"}).AddRow("Test User", "testemail@mail.com", "1234567890", "M")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM IPD_INFO_PERSONAL_DTLS, IAI_INFO_ACCOUNT_INFO, IAD_INFO_ADDRESS_DTLS WHERE (.+)$").
		WillReturnRows(expectedOutput)

	expectedPANOutput := sqlmock.NewRows([]string{"PanNo"}).AddRow("*")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, IPD_INFO_PERSONAL_DTLS, IAD_INFO_ADDRESS_DTLS, IID_INFO_IDENTIFICATION_DTLS, UAC_USR_ACCNTS WHERE (.+)$").
		WillReturnRows(expectedPANOutput)

	// Triggering Function
	_, err := suite.userStore.GetUserDetails(suite.ctx, "TestUser", "12345678", "Y")

	// Validations
	suite.EqualError(err, "pan not found")
}

func (suite *UserSuite) TestGetUserDetails_ReturnPrivacyInfoError() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"FullName", "Email", "MobileNo", "Gender"}).AddRow("Test User", "testemail@mail.com", "1234567890", "M")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM IPD_INFO_PERSONAL_DTLS, IAI_INFO_ACCOUNT_INFO, IAD_INFO_ADDRESS_DTLS WHERE (.+)$").
		WillReturnRows(expectedOutput)

	expectedPANOutput := sqlmock.NewRows([]string{"PanNo"}).AddRow("AAAAA0000A")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, IPD_INFO_PERSONAL_DTLS, IAD_INFO_ADDRESS_DTLS, IID_INFO_IDENTIFICATION_DTLS, UAC_USR_ACCNTS WHERE (.+)$").
		WillReturnRows(expectedPANOutput)

	suite.sqlMock.
		ExpectQuery("^SELECT (.+) CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnError(errors.New("ORA Error"))

	// Triggering Function
	_, err := suite.userStore.GetUserDetails(suite.ctx, "TestUser", "12345678", "Y")

	// Validations
	suite.EqualError(err, "ORA Error")
}

func (suite *UserSuite) TestGetUserDetails_ReturnOnlyPersonalInfo() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"FullName", "Email", "MobileNo", "Gender"}).AddRow("Test User", "testemail@mail.com", "1234567890", "M")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM IPD_INFO_PERSONAL_DTLS, IAI_INFO_ACCOUNT_INFO, IAD_INFO_ADDRESS_DTLS WHERE (.+)$").
		WillReturnRows(expectedOutput)
	expectedExtraInfoOutput := sqlmock.NewRows([]string{"PanNo", "DateOfBirth", "FirstName", "IsRICustomer"}).AddRow("AAAAA0000A", "01-Jan-1990", "Test User", "Y")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, IPD_INFO_PERSONAL_DTLS, IAD_INFO_ADDRESS_DTLS, IID_INFO_IDENTIFICATION_DTLS, UAC_USR_ACCNTS WHERE (.+)$").
		WillReturnRows(expectedExtraInfoOutput)

	// Triggering Function
	personalInfo, err := suite.userStore.GetUserDetails(suite.ctx, "TestUser", "12345678", "N")

	// Validations
	suite.NoError(err)
	suite.Equal("Test User", personalInfo.FullName)
	suite.Equal("testemail@mail.com", personalInfo.Email)
	suite.Equal("M", personalInfo.Gender)
	suite.Equal("1234567890", personalInfo.MobileNo)
	suite.Equal("Y", personalInfo.CustomerInfo.IsRICustomer)

	// Customer/Extra user info should not be available
	suite.Equal("", personalInfo.CustomerInfo.PanNo)
	suite.Equal("", personalInfo.ExtraPersonalInfo.FirstName)
}

func (suite *UserSuite) TestGetUserDetails_ReturnOnlyPersonalAndCustomerInfo_WhenPrivacyIsOpted() {
	// Mocking and Setting Expected Result
	expectedPersonalInfoOutput := sqlmock.NewRows([]string{"FullName", "Email", "MobileNo", "Gender"}).AddRow("Test User", "testemail@mail.com", "1234567890", "M")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM IPD_INFO_PERSONAL_DTLS, IAI_INFO_ACCOUNT_INFO, IAD_INFO_ADDRESS_DTLS WHERE (.+)$").
		WillReturnRows(expectedPersonalInfoOutput)

	expectedExtraInfoOutput := sqlmock.NewRows([]string{"PanNo", "DateOfBirth", "FirstName", "IsRICustomer"}).AddRow("AAAAA0000A", "01-Jan-1990", "Test User", "Y")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, IPD_INFO_PERSONAL_DTLS, IAD_INFO_ADDRESS_DTLS, IID_INFO_IDENTIFICATION_DTLS, UAC_USR_ACCNTS WHERE (.+)$").
		WillReturnRows(expectedExtraInfoOutput)

	expectedNodeId := sqlmock.NewRows([]string{"nodeID"}).AddRow("123456")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnRows(expectedNodeId)

	expectedTransactionCount := sqlmock.NewRows([]string{"transCount"}).AddRow("0")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) MF_TRANSACTIONS WHERE (.+)$").
		WillReturnRows(expectedTransactionCount)

	expectedNodeIDCount := sqlmock.NewRows([]string{"nodeCount"}).AddRow("0")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnRows(expectedNodeIDCount)

	// Triggering Function
	personalInfo, err := suite.userStore.GetUserDetails(suite.ctx, "TestUser", "12345678", "Y")

	// Validations
	suite.NoError(err)
	suite.Equal("Test User", personalInfo.FullName)
	suite.Equal("testemail@mail.com", personalInfo.Email)
	suite.Equal("M", personalInfo.Gender)
	suite.Equal("1234567890", personalInfo.MobileNo)

	suite.Equal("AAAAA0000A", personalInfo.CustomerInfo.PanNo)
	suite.Equal("Y", personalInfo.CustomerInfo.IsRICustomer)

	// Extra user info should not be available
	suite.Equal("", personalInfo.ExtraPersonalInfo.FirstName)
	suite.Equal("", personalInfo.ExtraPersonalInfo.DateOfBirth)
	suite.Equal("", personalInfo.ExtraPersonalInfo.EmailAddress)
}

func (suite *UserSuite) TestGetUserDetails_ReturnBothPersonalAndExtraInfo_WhenPrivacyIsNotOpted() {
	// Mocking and Setting Expected Result
	expectedPersonalInfoOutput := sqlmock.NewRows([]string{"FullName", "Email", "MobileNo", "Gender"}).AddRow("Test User", "testemail@mail.com", "1234567890", "M")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM IPD_INFO_PERSONAL_DTLS, IAI_INFO_ACCOUNT_INFO, IAD_INFO_ADDRESS_DTLS WHERE (.+)$").
		WillReturnRows(expectedPersonalInfoOutput)

	expectedExtraInfoOutput := sqlmock.NewRows([]string{"PanNo", "DateOfBirth", "FirstName", "LastName", "IsRICustomer"}).AddRow("AAAAA0000A", "01-Jan-1990", "Test", "User", "Y")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, IPD_INFO_PERSONAL_DTLS, IAD_INFO_ADDRESS_DTLS, IID_INFO_IDENTIFICATION_DTLS, UAC_USR_ACCNTS WHERE (.+)$").
		WillReturnRows(expectedExtraInfoOutput)

	expectedNodeId := sqlmock.NewRows([]string{"nodeID"}).AddRow("TestUser")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnRows(expectedNodeId)

	expectedTransactionCount := sqlmock.NewRows([]string{"transCount"}).AddRow("1")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) MF_TRANSACTIONS WHERE (.+)$").
		WillReturnRows(expectedTransactionCount)

	// Triggering Function
	personalInfo, err := suite.userStore.GetUserDetails(suite.ctx, "TestUser", "12345678", "Y")

	// Validations
	suite.NoError(err)
	suite.Equal("Test User", personalInfo.FullName)
	suite.Equal("testemail@mail.com", personalInfo.Email)
	suite.Equal("M", personalInfo.Gender)
	suite.Equal("1234567890", personalInfo.MobileNo)

	suite.Equal("AAAAA0000A", personalInfo.CustomerInfo.PanNo)
	suite.Equal("Y", personalInfo.CustomerInfo.IsRICustomer)

	// Extra user info should not be available
	suite.Equal("Test", personalInfo.ExtraPersonalInfo.FirstName)
	suite.Equal("User", personalInfo.ExtraPersonalInfo.LastName)
	suite.Equal("01-Jan-1990", personalInfo.ExtraPersonalInfo.DateOfBirth)
}

func (suite *UserSuite) TestGetPrivacyInfo_ReturnNodeIDSQLError() {
	// Mocking and Setting Expected Result
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnError(errors.New("ORA Error"))

	// Triggering Function
	_, err := suite.userStore.GetPrivacyInfo(suite.ctx, "TestUser", "AAAAA0000A", false)

	// Validations
	suite.EqualError(err, "ORA Error")
}

func (suite *UserSuite) TestGetPrivacyInfo_ReturnNodeIDNotFoundError() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"nodeID"})
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnRows(expectedOutput)

	// Triggering Function
	_, err := suite.userStore.GetPrivacyInfo(suite.ctx, "TestUser", "AAAAA0000A", false)

	// Validations
	suite.EqualError(err, "no entry found for this pan in CMM table")
}

func (suite *UserSuite) TestGetPrivacyInfo_ReturnTransactionCountSQLError() {
	// Mocking and Setting Expected Result
	expectedOutput := sqlmock.NewRows([]string{"nodeID"}).AddRow("123456")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnRows(expectedOutput)
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) MF_TRANSACTIONS WHERE (.+)$").
		WillReturnError(errors.New("ORA Error"))

	// Triggering Function
	_, err := suite.userStore.GetPrivacyInfo(suite.ctx, "TestUser", "AAAAA0000A", false)

	// Validations
	suite.EqualError(err, "ORA Error")
}

func (suite *UserSuite) TestGetPrivacyInfo_ReturnOutput_WhenNodeIDIsSameAsUserID() {
	// Mocking and Setting Expected Result
	expectedNodeId := sqlmock.NewRows([]string{"nodeID"}).AddRow("TestUser")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnRows(expectedNodeId)

	expectedTransactionCount := sqlmock.NewRows([]string{"transCount"}).AddRow("0")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) MF_TRANSACTIONS WHERE (.+)$").
		WillReturnRows(expectedTransactionCount)

	// Triggering Function
	actualOutput, err := suite.userStore.GetPrivacyInfo(suite.ctx, "TestUser", "AAAAA0000A", false)

	// Validations
	suite.NoError(err)
	suite.Equal(true, actualOutput.PrivacyOpted)
	suite.Equal("N", actualOutput.ModifyAllowed)
}

func (suite *UserSuite) TestGetPrivacyInfo_ReturnOutput_WhenTransactionCountIsAboveZero() {
	// Mocking and Setting Expected Result
	expectedNodeId := sqlmock.NewRows([]string{"nodeID"}).AddRow("123456")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnRows(expectedNodeId)

	expectedTransactionCount := sqlmock.NewRows([]string{"transCount"}).AddRow("1")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) MF_TRANSACTIONS WHERE (.+)$").
		WillReturnRows(expectedTransactionCount)

	// Triggering Function
	actualOutput, err := suite.userStore.GetPrivacyInfo(suite.ctx, "TestUser", "AAAAA0000A", false)

	// Validations
	suite.NoError(err)
	suite.Equal(true, actualOutput.PrivacyOpted)
	suite.Equal("N", actualOutput.ModifyAllowed)
}

func (suite *UserSuite) TestGetPrivacyInfo_ReturnOutputWithModifyFlag() {
	// Mocking and Setting Expected Result
	expectedNodeId := sqlmock.NewRows([]string{"nodeID"}).AddRow("123456")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnRows(expectedNodeId)

	expectedTransactionCount := sqlmock.NewRows([]string{"transCount"}).AddRow("1")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) MF_TRANSACTIONS WHERE (.+)$").
		WillReturnRows(expectedTransactionCount)

	// Triggering Function
	actualOutput, err := suite.userStore.GetPrivacyInfo(suite.ctx, "TestUser", "AAAAA0000A", true)

	// Validations
	suite.NoError(err)
	suite.Equal(true, actualOutput.PrivacyOpted)
	suite.Equal("U", actualOutput.ModifyAllowed)
}

func (suite *UserSuite) TestGetPrivacyInfo_ReturnNodeIDCountSQLError() {
	// Mocking and Setting Expected Result
	expectedNodeId := sqlmock.NewRows([]string{"nodeID"}).AddRow("123456")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnRows(expectedNodeId)

	expectedTransactionCount := sqlmock.NewRows([]string{"transCount"}).AddRow("0")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) MF_TRANSACTIONS WHERE (.+)$").
		WillReturnRows(expectedTransactionCount)

	suite.sqlMock.
		ExpectQuery("^SELECT (.+) CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnError(errors.New("ORA Error"))

	// Triggering Function
	_, err := suite.userStore.GetPrivacyInfo(suite.ctx, "TestUser", "AAAAA0000A", false)

	// Validations
	suite.EqualError(err, "ORA Error")
}

func (suite *UserSuite) TestGetPrivacyInfo_ReturnOutput_WhenNodeIDCountIsZero() {
	// Mocking and Setting Expected Result
	expectedNodeId := sqlmock.NewRows([]string{"nodeID"}).AddRow("123456")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnRows(expectedNodeId)

	expectedTransactionCount := sqlmock.NewRows([]string{"transCount"}).AddRow("0")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) MF_TRANSACTIONS WHERE (.+)$").
		WillReturnRows(expectedTransactionCount)

	expectedNodeIDCount := sqlmock.NewRows([]string{"nodeCount"}).AddRow("0")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnRows(expectedNodeIDCount)

	// Triggering Function
	actualOutput, err := suite.userStore.GetPrivacyInfo(suite.ctx, "TestUser", "AAAAA0000A", false)

	// Validations
	suite.NoError(err)
	suite.Equal(false, actualOutput.PrivacyOpted)
	suite.Equal("M", actualOutput.ModifyAllowed)
}

func (suite *UserSuite) TestGetPrivacyInfo_ReturnOutput_WhenNodeIDCountIsAboveZero() {
	// Mocking and Setting Expected Result
	expectedNodeId := sqlmock.NewRows([]string{"nodeID"}).AddRow("123456")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnRows(expectedNodeId)

	expectedTransactionCount := sqlmock.NewRows([]string{"transCount"}).AddRow("0")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) MF_TRANSACTIONS WHERE (.+)$").
		WillReturnRows(expectedTransactionCount)

	expectedNodeIDCount := sqlmock.NewRows([]string{"nodeCount"}).AddRow("1")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) CMM_CLNT_MAP_MSTR WHERE (.+)$").
		WillReturnRows(expectedNodeIDCount)

	// Triggering Function
	actualOutput, err := suite.userStore.GetPrivacyInfo(suite.ctx, "TestUser", "AAAAA0000A", false)

	// Validations
	suite.NoError(err)
	suite.Equal(false, actualOutput.PrivacyOpted)
	suite.Equal("N", actualOutput.ModifyAllowed)
}

func (suite *UserSuite) TestGetEuinNo_CallAndTradeUser_ReturnSQLError() {
	// Mocking and Setting Expected Result
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM BPA_EMP_EUIN_DTLS, HEE_HR_EMP_EXTRCT WHERE (.+)$").
		WillReturnError(errors.New("ORA Error"))

	// Triggering Function
	_, err := suite.userStore.GetEuinNo(suite.ctx, "TestUser", "12345678", "700000", false)

	// Validations
	suite.Error(err, "ORA Error")
}

func (suite *UserSuite) TestGetEuinNo_CallAndTradeUser_ReturnValidEuinNo_WhenNoRecordsFound() {
	// Mocking and Setting Expected Result
	mockedRow := sqlmock.NewRows([]string{"euinNo"})
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM BPA_EMP_EUIN_DTLS, HEE_HR_EMP_EXTRCT WHERE (.+)$").
		WillReturnRows(mockedRow)

	// Triggering Function
	euinNo, err := suite.userStore.GetEuinNo(suite.ctx, "TestUser", "12345678", "700000", false)

	suite.NoError(err)
	suite.Equal(euinNo, "")
}

func (suite *UserSuite) TestGetEuinNo_CallAndTradeUser_ReturnValidEuinNo_WhenRecordsFound() {
	// Mocking and Setting Expected Result
	mockedRow := sqlmock.NewRows([]string{"euinNo"}).AddRow("EUIN123456")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM BPA_EMP_EUIN_DTLS, HEE_HR_EMP_EXTRCT WHERE (.+)$").
		WillReturnRows(mockedRow)

	// Triggering Function
	euinNo, err := suite.userStore.GetEuinNo(suite.ctx, "TestUser", "12345678", "700000", false)

	suite.NoError(err)
	suite.Equal(euinNo, "EUIN123456")
}

func (suite *UserSuite) TestGetEuinNo_BusinessPartner_ReturnSQLError_WhileFetchingFamilyType() {
	// Mocking and Setting Expected Result
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICAD_INFO_CLIENT_ADDL_DTLS, IAI_INFO_ACCOUNT_INFO WHERE (.+)$").
		WillReturnError(errors.New("ORA Error"))

	// Triggering Function
	_, err := suite.userStore.GetEuinNo(suite.ctx, "#BPID", "12345678", "", false)

	// Validations
	suite.Error(err, "ORA Error")
}

func (suite *UserSuite) TestGetEuinNo_BusinessPartner_ReturnEuinAsEmpty_ForFamilyAccount() {
	// Mocking and Setting Expected Result
	familyTypeMockedRow := sqlmock.NewRows([]string{"familyType"}).AddRow("FAN")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICAD_INFO_CLIENT_ADDL_DTLS, IAI_INFO_ACCOUNT_INFO WHERE (.+)$").
		WillReturnRows(familyTypeMockedRow)

	// Triggering Function
	euinNo, err := suite.userStore.GetEuinNo(suite.ctx, "#BPID", "12345678", "", false)

	// Validations
	suite.NoError(err)
	suite.Equal(euinNo, "")
}

func (suite *UserSuite) TestGetEuinNo_BusinessPartner_ReturnSQLError_WhileFetchingEuinNo() {
	// Mocking and Setting Expected Result
	familyTypeMockedRow := sqlmock.NewRows([]string{"familyType"}).AddRow("NA")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICAD_INFO_CLIENT_ADDL_DTLS, IAI_INFO_ACCOUNT_INFO WHERE (.+)$").
		WillReturnRows(familyTypeMockedRow)
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM BPA_CTCL_DETAILS WHERE (.+)$").
		WillReturnError(errors.New("ORA Error"))

	// Triggering Function
	_, err := suite.userStore.GetEuinNo(suite.ctx, "#BPID", "12345678", "", false)

	// Validations
	suite.Error(err, "ORA Error")
}

func (suite *UserSuite) TestGetEuinNo_BusinessPartner_ReturnEuinNoAsEmpty_WhenNoEuinFound() {
	// Mocking and Setting Expected Result
	familyTypeMockedRow := sqlmock.NewRows([]string{"familyType"}).AddRow("NA")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICAD_INFO_CLIENT_ADDL_DTLS, IAI_INFO_ACCOUNT_INFO WHERE (.+)$").
		WillReturnRows(familyTypeMockedRow)
	euinNoMockedRow := sqlmock.NewRows([]string{"euinNo"})
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM BPA_CTCL_DETAILS WHERE (.+)$").
		WillReturnRows(euinNoMockedRow)

	// Triggering Function
	euinNo, err := suite.userStore.GetEuinNo(suite.ctx, "#BPID", "12345678", "", false)

	// Validations
	suite.NoError(err)
	suite.Equal(euinNo, "")
}

func (suite *UserSuite) TestGetEuinNo_BusinessPartner_ReturnValidEuinNo_WhenEuinFound() {
	// Mocking and Setting Expected Result
	familyTypeMockedRow := sqlmock.NewRows([]string{"familyType"}).AddRow("NA")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM ICAD_INFO_CLIENT_ADDL_DTLS, IAI_INFO_ACCOUNT_INFO WHERE (.+)$").
		WillReturnRows(familyTypeMockedRow)
	euinNoMockedRow := sqlmock.NewRows([]string{"euinNo"}).AddRow("EUIN123456")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM BPA_CTCL_DETAILS WHERE (.+)$").
		WillReturnRows(euinNoMockedRow)

	// Triggering Function
	euinNo, err := suite.userStore.GetEuinNo(suite.ctx, "#BPID", "12345678", "", false)

	// Validations
	suite.NoError(err)
	suite.Equal(euinNo, "EUIN123456")
}

func (suite *UserSuite) TestGetEuinNo_RegularUser_ReturnEuinNoAsEmpty_ForD2uUsers() {
	// Triggering Function
	euinNo, err := suite.userStore.GetEuinNo(suite.ctx, "TestUser", "12345678", "", true)

	// Validations
	suite.NoError(err)
	suite.Equal(euinNo, "")
}

func (suite *UserSuite) TestGetEuinNo_RegularUser_ReturnSQLError_WhileFetchingBPID() {
	// Mocking and Setting Expected Result
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CLM_CLNT_MSTR WHERE (.+)$").
		WillReturnError(errors.New("ORA Error"))

	// Triggering Function
	_, err := suite.userStore.GetEuinNo(suite.ctx, "TestUser", "12345678", "", false)

	// Validations
	suite.Error(err, "ORA Error")
}

func (suite *UserSuite) TestGetEuinNo_RegularUser_ReturnEuinNoAsEmpty_WhenBPIDNotFound() {
	// Mocking and Setting Expected Result
	bpIdMockedRow := sqlmock.NewRows([]string{"bpID"}).AddRow("N")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CLM_CLNT_MSTR WHERE (.+)$").
		WillReturnRows(bpIdMockedRow)

	// Triggering Function
	euinNo, err := suite.userStore.GetEuinNo(suite.ctx, "TestUser", "12345678", "", false)

	// Validations
	suite.NoError(err)
	suite.Equal(euinNo, "")
}

func (suite *UserSuite) TestGetEuinNo_RegularUser_ReturnEuinNoAsEmpty_WhenEuinSQLError() {
	// Mocking and Setting Expected Result
	bpIdMockedRow := sqlmock.NewRows([]string{"bpID"}).AddRow("#BPID")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CLM_CLNT_MSTR WHERE (.+)$").
		WillReturnRows(bpIdMockedRow)
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM BPA_CTCL_DETAILS WHERE (.+)$").
		WillReturnError(errors.New("ORA Error"))

	// Triggering Function
	_, err := suite.userStore.GetEuinNo(suite.ctx, "TestUser", "12345678", "", false)

	// Validations
	suite.Error(err, "ORA Error")
}

func (suite *UserSuite) TestGetEuinNo_RegularUser_ReturnValidEuinNo_WhenBPIDFound() {
	// Mocking and Setting Expected Result
	bpIdMockedRow := sqlmock.NewRows([]string{"bpID"}).AddRow("#BPID")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM CLM_CLNT_MSTR WHERE (.+)$").
		WillReturnRows(bpIdMockedRow)
	euinNoMockedRow := sqlmock.NewRows([]string{"euinNo"}).AddRow("EUIN123456")
	suite.sqlMock.
		ExpectQuery("^SELECT (.+) FROM BPA_CTCL_DETAILS WHERE (.+)$").
		WillReturnRows(euinNoMockedRow)

	// Triggering Function
	euinNo, err := suite.userStore.GetEuinNo(suite.ctx, "TestUser", "12345678", "", false)

	// Validations
	suite.NoError(err)
	suite.Equal(euinNo, "EUIN123456")
}

func (suite *UserSuite) TestGetSipHealthFlag() {
	testCases := []struct {
		desc           string
		mockInput      *sqlmock.Rows
		expectedError  string
		expectedOutput string
	}{
		{
			desc:           "SQLError",
			mockInput:      nil,
			expectedError:  "ORA Error",
			expectedOutput: "",
		}, {
			desc:           "NoRowsError",
			mockInput:      sqlmock.NewRows([]string{"sipHealthFlag"}),
			expectedError:  "",
			expectedOutput: "N",
		}, {
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"sipHealthFlag"}).AddRow("Y"),
			expectedError:  "",
			expectedOutput: "Y",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM CSD_CLNT_SPT_DTLS WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM CSD_CLNT_SPT_DTLS WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.userStore.GetSipHealthFlag(suite.ctx, "12345678")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput, testCase.expectedOutput)
			}

		})
	}
}

func (suite *UserSuite) TestGetEBAUploadDate() {
	testCases := []struct {
		desc           string
		mockInput      *sqlmock.Rows
		expectedError  string
		expectedOutput string
	}{
		{
			desc:           "SQLError",
			mockInput:      nil,
			expectedError:  "ORA Error",
			expectedOutput: "",
		}, {
			desc:           "NoRowsError",
			mockInput:      sqlmock.NewRows([]string{"ebaUploadDate"}),
			expectedError:  "",
			expectedOutput: "",
		}, {
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"ebaUploadDate"}).AddRow("25/01/2025"),
			expectedError:  "",
			expectedOutput: "25/01/2025",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, ICAD_INFO_CLIENT_ADDL_DTLS WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM ICD_INFO_CLIENT_DTLS, ICAD_INFO_CLIENT_ADDL_DTLS WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.userStore.GetEBAUploadDate(suite.ctx, "test")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput, testCase.expectedOutput)
			}

		})
	}
}

func (suite *UserSuite) TestIsD2UEnabled() {
	testCases := []struct {
		desc           string
		mockInput      *sqlmock.Rows
		expectedError  string
		expectedOutput bool
	}{
		{
			desc:           "SQLError",
			mockInput:      nil,
			expectedError:  "ORA Error",
			expectedOutput: false,
		}, {
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"rowCount"}).AddRow(1),
			expectedError:  "",
			expectedOutput: true,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM DMM_D2U_MATCH_MPPNG_MSTR WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM DMM_D2U_MATCH_MPPNG_MSTR WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.userStore.IsD2UEnabled(suite.ctx, "12345678")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput, testCase.expectedOutput)
			}

		})
	}
}

func (suite *UserSuite) TestIsD2UActive() {
	testCases := []struct {
		desc           string
		mockInput      *sqlmock.Rows
		expectedError  string
		expectedOutput bool
	}{
		{
			desc:           "SQLError",
			mockInput:      nil,
			expectedError:  "ORA Error",
			expectedOutput: false,
		}, {
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"rowCount"}).AddRow(1),
			expectedError:  "",
			expectedOutput: true,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM DCM_D2U_CLNT_MSTR, UAC_USR_ACCNTS WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM DCM_D2U_CLNT_MSTR, UAC_USR_ACCNTS WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.userStore.IsD2UActive(suite.ctx, "12345678")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput, testCase.expectedOutput)
			}

		})
	}
}

func (suite *UserSuite) TestIsEATMEnabled() {
	testCases := []struct {
		desc           string
		mockInput      *sqlmock.Rows
		expectedError  string
		expectedOutput bool
	}{
		{
			desc:          "SQLError",
			mockInput:     nil,
			expectedError: "ORA Error",
		}, {
			desc:           "NoRowsError",
			mockInput:      sqlmock.NewRows([]string{"eATMFlag"}),
			expectedOutput: false,
		}, {
			desc:           "SuccessReturnsFalse",
			mockInput:      sqlmock.NewRows([]string{"eATMFlag"}).AddRow("N"),
			expectedOutput: false,
		}, {
			desc:           "SuccessReturnsTrue",
			mockInput:      sqlmock.NewRows([]string{"eATMFlag"}).AddRow("Y"),
			expectedOutput: true,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM CSD_CLNT_SPT_DTLS  WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM CSD_CLNT_SPT_DTLS  WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.userStore.IsEATMEnabled(suite.ctx, "12345678")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedOutput, actualOutput)
			}
		})
	}
}

func (suite *UserSuite) TestGetAccountType() {
	testCases := []struct {
		desc           string
		mockInput      *sqlmock.Rows
		expectedError  string
		expectedOutput string
	}{
		{
			desc:          "SQLError",
			mockInput:     nil,
			expectedError: "ORA Error",
		}, {
			desc:          "NoRowsError",
			mockInput:     sqlmock.NewRows([]string{"accountType"}),
			expectedError: "No user found",
		}, {
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"accountType"}).AddRow("O"),
			expectedOutput: "O",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM IAI_INFO_ACCOUNT_INFO  WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM IAI_INFO_ACCOUNT_INFO  WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.userStore.GetAccountType(suite.ctx, "12345678")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedOutput, actualOutput)
			}
		})
	}
}
