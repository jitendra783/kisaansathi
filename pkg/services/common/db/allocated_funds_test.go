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

type AllocatedFundsSuite struct {
	suite.Suite
	ctx                      context.Context
	sqlDB                    *sql.DB
	gormDB                   *gorm.DB
	sqlMock                  sqlmock.Sqlmock
	allocatedFundsSuiteStore AllocatedFundsStore
}

func TestAllocatedFundsSuite(t *testing.T) {
	suite.Run(t, new(AllocatedFundsSuite))
}

func (suite *AllocatedFundsSuite) SetupSuite() {
	logger.LoggerInit("", -1)

	suite.ctx = context.TODO()
	suite.sqlDB, suite.gormDB, suite.sqlMock = utils.NewMockDB()
	suite.allocatedFundsSuiteStore = NewAllocatedFundsStore(suite.gormDB)
}

func (suite *AllocatedFundsSuite) TestGetLinkedLimits() {
	testCases := []struct {
		desc                string
		mockInput           *sqlmock.Rows
		expectedError       string
		expectedLimitAmount float64
	}{
		{
			desc:                "SQLError",
			mockInput:           nil,
			expectedError:       "ORA Error",
			expectedLimitAmount: 0.00,
		}, {
			desc:                "ReturnsAmountForCreditFlow",
			mockInput:           sqlmock.NewRows([]string{"allocatedAmount", "netBalance", "debitCreditFlag"}).AddRow(150.0, 50.0, "C"),
			expectedError:       "",
			expectedLimitAmount: 150.00,
		}, {
			desc:                "ReturnsAmountForDebitFlow",
			mockInput:           sqlmock.NewRows([]string{"allocatedAmount", "netBalance", "debitCreditFlag"}).AddRow(150.0, 50.0, "D"),
			expectedError:       "",
			expectedLimitAmount: 100.00,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM CLM_CLNT_MSTR, MF_TOT_BLNCS WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM CLM_CLNT_MSTR, MF_TOT_BLNCS WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.allocatedFundsSuiteStore.GetLinkedLimits(suite.ctx, "12345678")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput, testCase.expectedLimitAmount)
			}
		})
	}
}

func (suite *AllocatedFundsSuite) TestGetDepositLinkedLimits() {
	testCases := []struct {
		desc                string
		mockInput           *sqlmock.Rows
		expectedError       string
		expectedLimitAmount float64
	}{
		{
			desc:                "SQLError",
			mockInput:           nil,
			expectedError:       "ORA Error",
			expectedLimitAmount: 0.00,
		}, {
			desc:                "ReturnsAmountForCreditFlow",
			mockInput:           sqlmock.NewRows([]string{"allocatedAmount", "netBalance", "debitCreditFlag"}).AddRow(150.0, 50.0, "C"),
			expectedError:       "",
			expectedLimitAmount: 150.00,
		}, {
			desc:                "ReturnsAmountForDebitFlow",
			mockInput:           sqlmock.NewRows([]string{"allocatedAmount", "netBalance", "debitCreditFlag"}).AddRow(150.0, 50.0, "D"),
			expectedError:       "",
			expectedLimitAmount: 100.00,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM CLM_CLNT_MSTR, MF_TOT_BLNCS WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM CLM_CLNT_MSTR, MF_TOT_BLNCS WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.allocatedFundsSuiteStore.GetDepositLinkedLimits(suite.ctx, "12345678")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput, testCase.expectedLimitAmount)
			}
		})
	}
}

func (suite *AllocatedFundsSuite) TestGetTPALimits() {
	testCases := []struct {
		desc                string
		mockInput           *sqlmock.Rows
		expectedError       string
		expectedLimitAmount float64
	}{
		{
			desc:                "SQLError",
			mockInput:           nil,
			expectedError:       "ORA Error",
			expectedLimitAmount: 0.00,
		}, {
			desc:                "ReturnsAmountForCreditFlow",
			mockInput:           sqlmock.NewRows([]string{"allocatedAmount", "netBalance", "debitCreditFlag"}).AddRow(150.0, 50.0, "C"),
			expectedError:       "",
			expectedLimitAmount: 150.00,
		}, {
			desc:                "ReturnsAmountForDebitFlow",
			mockInput:           sqlmock.NewRows([]string{"allocatedAmount", "netBalance", "debitCreditFlag"}).AddRow(150.0, 50.0, "D"),
			expectedError:       "",
			expectedLimitAmount: 100.00,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM CLM_CLNT_MSTR, MTB_MF_TPA_BLNCS WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM CLM_CLNT_MSTR, MTB_MF_TPA_BLNCS WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.allocatedFundsSuiteStore.GetTPALimits(suite.ctx, "12345678")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput, testCase.expectedLimitAmount)
			}
		})
	}
}
