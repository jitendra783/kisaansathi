package db

import (
	"context"
	"database/sql"
	"errors"
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/common/models"
	"kisaanSathi/pkg/utils"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type SchemeSuite struct {
	suite.Suite
	ctx         context.Context
	sqlDB       *sql.DB
	gormDB      *gorm.DB
	sqlMock     sqlmock.Sqlmock
	schemeStore SchemeStore
}

func TestSchemeSuite(t *testing.T) {
	suite.Run(t, new(SchemeSuite))
}

func (suite *SchemeSuite) SetupSuite() {
	logger.LoggerInit("", -1)

	suite.ctx = context.TODO()
	suite.sqlDB, suite.gormDB, suite.sqlMock = utils.NewMockDB()
	suite.schemeStore = NewSchemeStore(suite.gormDB)
}

func (suite *SchemeSuite) TestGetCompanyRegistrar() {
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
			mockInput:     sqlmock.NewRows([]string{"registrar"}),
			expectedError: "No registrar found for the company",
		}, {
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"registrar"}).AddRow("KARVY"),
			expectedOutput: "KARVY",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_COMPANIES WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_COMPANIES WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.GetCompanyRegistrar(suite.ctx, 19)

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

func (suite *SchemeSuite) TestGetSchemeFlags() {
	testCases := []struct {
		desc           string
		mockInput      *sqlmock.Rows
		expectedError  string
		expectedOutput *models.SchemeFlags
	}{
		{
			desc:          "SQLError",
			mockInput:     nil,
			expectedError: "ORA Error",
		}, {
			desc:          "NoRowsError",
			mockInput:     sqlmock.NewRows([]string{"CloseFlag", "DirectSchemeFlag", "PurchaseFlag"}),
			expectedError: "No scheme found",
		}, {
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"CloseFlag", "DirectSchemeFlag", "PurchaseFlag"}).AddRow("Open Ended", "N", "Y"),
			expectedOutput: &models.SchemeFlags{CloseFlag: "Open Ended", PurchaseFlag: "Y", DirectSchemeFlag: "N"},
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.GetSchemeFlags(suite.ctx, 19, "IOGP")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedOutput.CloseFlag, actualOutput.CloseFlag)
				assert.Equal(t, testCase.expectedOutput.DirectSchemeFlag, actualOutput.DirectSchemeFlag)
				assert.Equal(t, testCase.expectedOutput.PurchaseFlag, actualOutput.PurchaseFlag)
				assert.Equal(t, testCase.expectedOutput.SIPFlag, actualOutput.SIPFlag)
			}
		})
	}
}

func (suite *SchemeSuite) TestGetSchemeDetails() {
	testCases := []struct {
		desc           string
		mockInput      *sqlmock.Rows
		expectedError  string
		expectedOutput *models.SchemeDetails
	}{
		{
			desc:          "SQLError",
			mockInput:     nil,
			expectedError: "ORA Error",
		}, {
			desc:          "NoRowsError",
			mockInput:     sqlmock.NewRows([]string{"MinPurchaseAmount", "MultiPurchaseAmount", "SchemeDesc"}),
			expectedError: "No scheme found",
		}, {
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"MinPurchaseAmount", "MultiPurchaseAmount", "SchemeDesc"}).AddRow(100, 10, "Test Scheme"),
			expectedOutput: &models.SchemeDetails{MinPurchaseAmount: 100, MultiPurchaseAmount: 10, SchemeDesc: "Test Scheme"},
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.GetSchemeDetails(suite.ctx, 19, "IOGP")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedOutput.MinPurchaseAmount, actualOutput.MinPurchaseAmount)
				assert.Equal(t, testCase.expectedOutput.MultiPurchaseAmount, actualOutput.MultiPurchaseAmount)
				assert.Equal(t, testCase.expectedOutput.SchemeDesc, actualOutput.SchemeDesc)
			}
		})
	}
}

func (suite *SchemeSuite) TestGetSchemeNavDetails() {
	testCases := []struct {
		desc           string
		mockInput      *sqlmock.Rows
		expectedError  string
		expectedOutput *models.NavDetails
	}{
		{
			desc:          "SQLError",
			mockInput:     nil,
			expectedError: "ORA Error",
		}, {
			desc:          "NoRowsError",
			mockInput:     sqlmock.NewRows([]string{"AmfiCode", "NavDate", "NavValue"}),
			expectedError: "No nav found",
		}, {
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"AmfiCode", "NavDate", "NavValue"}).AddRow("1234567", "10-Jan-2025", 128.56),
			expectedOutput: &models.NavDetails{AmfiCode: "1234567", NavDate: "10-Jan-2025", NavValue: 128.56},
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_NAVS WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_NAVS WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.GetSchemeNavDetails(suite.ctx, 19, "IOGP")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedOutput.AmfiCode, actualOutput.AmfiCode)
				assert.Equal(t, testCase.expectedOutput.NavDate, actualOutput.NavDate)
				assert.Equal(t, testCase.expectedOutput.NavValue, actualOutput.NavValue)
			}
		})
	}
}

func (suite *SchemeSuite) TestGetCompanyFolio() {
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
			mockInput:     sqlmock.NewRows([]string{"compFolioFlag"}),
			expectedError: "No company found",
		}, {
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"compFolioFlag"}).AddRow("Y"),
			expectedOutput: "Y",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_COMPANIES WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_COMPANIES WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.GetCompanyFolio(suite.ctx, 19)

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

func (suite *SchemeSuite) TestGetRedeemSchemeDetails() {
	testCases := []struct {
		desc           string
		mockInput      *sqlmock.Rows
		expectedError  string
		expectedOutput *models.RedeemSchemeDetails
	}{
		{
			desc:          "SQLError",
			mockInput:     nil,
			expectedError: "ORA Error",
		}, {
			desc:          "NoRowsError",
			mockInput:     sqlmock.NewRows([]string{"AmfiCode", "NavDate", "NavValue", "AccountTypeAllowedFlag"}),
			expectedError: "No scheme found",
		}, {
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"AmfiCode", "NavDate", "NavValue", "AccountTypeAllowedFlag"}).AddRow("1234567", "11-Jan-2025", 101.56, "Y"),
			expectedOutput: &models.RedeemSchemeDetails{AmfiCode: "1234567", NavValue: 101.56, NavDate: "11-Jan-2025", AccountTypeAllowedFlag: "Y"},
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_NAVS, MF_SCHEME_MASTER WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_NAVS, MF_SCHEME_MASTER WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.GetRedeemSchemeDetails(suite.ctx, 19, "IOGP", "D", "O")

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

func (suite *SchemeSuite) TestIsParamEATMEnabled() {
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
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"eATMFlag"}).AddRow("Y"),
			expectedOutput: true,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_PARAM$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_PARAM$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.IsParamEATMEnabled(suite.ctx)

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

func (suite *SchemeSuite) TestIsSchemeEATMEnabled() {
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
			mockInput:      sqlmock.NewRows([]string{"compEATMFlag", "schemeEATMFlag", "categoryEATMFlag"}),
			expectedOutput: false,
		}, {
			desc:           "SuccessReturnsFalse",
			mockInput:      sqlmock.NewRows([]string{"compEATMFlag", "schemeEATMFlag", "categoryEATMFlag"}).AddRow("Y", "N", "Y"),
			expectedOutput: false,
		}, {
			desc:           "SuccessReturnsTrue",
			mockInput:      sqlmock.NewRows([]string{"compEATMFlag", "schemeEATMFlag", "categoryEATMFlag"}).AddRow("Y", "Y", "Y"),
			expectedOutput: true,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_COMPANIES, MF_SCHEME_MASTER, MF_CATEGORY_MASTER  WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_COMPANIES, MF_SCHEME_MASTER, MF_CATEGORY_MASTER  WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.IsSchemeEATMEnabled(suite.ctx, 19, "IOGP")

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
