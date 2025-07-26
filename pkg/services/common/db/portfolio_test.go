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

type PortfolioSuite struct {
	suite.Suite
	ctx            context.Context
	sqlDB          *sql.DB
	gormDB         *gorm.DB
	sqlMock        sqlmock.Sqlmock
	portfolioStore PortfolioStore
}

func TestPortfolioSuite(t *testing.T) {
	suite.Run(t, new(PortfolioSuite))
}

func (suite *PortfolioSuite) SetupSuite() {
	logger.LoggerInit("", -1)

	suite.ctx = context.TODO()
	suite.sqlDB, suite.gormDB, suite.sqlMock = utils.NewMockDB()
	suite.portfolioStore = NewPortfolioStore(suite.gormDB)
}

func (suite *PortfolioSuite) TestGetDPDetails() {
	testCases := []struct {
		desc                 string
		mockInput            *sqlmock.Rows
		expectedError        string
		expectedDpID         string
		expectedDpAccountNo  string
		expectedReinvestFlag string
	}{
		{
			desc:                 "SQLError",
			mockInput:            nil,
			expectedError:        "ORA Error",
			expectedDpID:         "",
			expectedDpAccountNo:  "",
			expectedReinvestFlag: "",
		}, {
			desc:                 "NoRowsError",
			mockInput:            sqlmock.NewRows([]string{"DpID", "DpAccountNo", "ReinvestFlag"}),
			expectedError:        "",
			expectedDpID:         "",
			expectedDpAccountNo:  "",
			expectedReinvestFlag: "X",
		}, {
			desc:                 "Success",
			mockInput:            sqlmock.NewRows([]string{"DpID", "DpAccountNo", "ReinvestFlag"}).AddRow("DP", "12334", "Y"),
			expectedError:        "",
			expectedDpID:         "DP",
			expectedDpAccountNo:  "12334",
			expectedReinvestFlag: "Y",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_UNIT_BAL WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_UNIT_BAL WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.portfolioStore.GetDPDetails(suite.ctx, 19, "IOGP", "7654321", "12345678")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedDpID, actualOutput.DpID)
				assert.Equal(t, testCase.expectedDpAccountNo, actualOutput.DpAccountNo)
				assert.Equal(t, testCase.expectedReinvestFlag, actualOutput.ReinvestFlag)
			}
		})
	}
}

func (suite *PortfolioSuite) TestIsTransactionExistsForScheme() {
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
			desc:           "NoRowsError",
			mockInput:      sqlmock.NewRows([]string{"transCount"}),
			expectedError:  "",
			expectedOutput: false,
		}, {
			desc:           "SuccessWithZeroRecords",
			mockInput:      sqlmock.NewRows([]string{"transCount"}).AddRow(0),
			expectedError:  "",
			expectedOutput: false,
		}, {
			desc:           "Success",
			mockInput:      sqlmock.NewRows([]string{"transCount"}).AddRow(1),
			expectedError:  "",
			expectedOutput: true,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_UNIT_BAL WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_UNIT_BAL WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.portfolioStore.IsTransactionExistsForScheme(suite.ctx, 19, "IOGP", "12345678")

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

func (suite *PortfolioSuite) TestGetDematUnblockedUnits() {
	testCases := []struct {
		desc                 string
		mockInput            *sqlmock.Rows
		expectedError        string
		expectedNoOfUnits    float64
		expectedReinvestFlag string
	}{
		{
			desc:                 "SQLError",
			mockInput:            nil,
			expectedError:        "ORA Error",
			expectedNoOfUnits:    0,
			expectedReinvestFlag: "",
		}, {
			desc:                 "NoRowsError",
			mockInput:            sqlmock.NewRows([]string{"NoOfUnits"}),
			expectedError:        "",
			expectedNoOfUnits:    0,
			expectedReinvestFlag: "X",
		}, {
			desc:                 "Success",
			mockInput:            sqlmock.NewRows([]string{"NoOfUnits"}).AddRow(10.545),
			expectedError:        "",
			expectedNoOfUnits:    10.545,
			expectedReinvestFlag: "",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MDBD_DP_BLCK_DTLS WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MDBD_DP_BLCK_DTLS WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.portfolioStore.GetDematUnblockedUnits(suite.ctx, "12345678", "DP_ID", "CLIENT_ID", "STOCK_CODE", "ISIN")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedNoOfUnits, actualOutput.NoOfUnits)
				assert.Equal(t, testCase.expectedReinvestFlag, actualOutput.ReinvestFlag)
			}
		})
	}
}
func (suite *PortfolioSuite) TestGetUnblockedUnits() {
	testCases := []struct {
		desc                      string
		mockInput                 *sqlmock.Rows
		expectedError             string
		expectedNoOfUnits         float64
		expectedReinvestFlag      string
		expectedFolioCreationDate string
	}{
		{
			desc:              "SQLError",
			mockInput:         nil,
			expectedError:     "ORA Error",
			expectedNoOfUnits: 0,
		}, {
			desc:                      "NoRowsError",
			mockInput:                 sqlmock.NewRows([]string{"NoOfUnits", "FolioCreationDate", "ReinvestFlag"}),
			expectedError:             "",
			expectedNoOfUnits:         0,
			expectedReinvestFlag:      "X",
			expectedFolioCreationDate: "",
		}, {
			desc:                      "Success",
			mockInput:                 sqlmock.NewRows([]string{"NoOfUnits", "FolioCreationDate", "ReinvestFlag"}).AddRow(12.545, "09-08-2024", "Y"),
			expectedError:             "",
			expectedNoOfUnits:         12.545,
			expectedReinvestFlag:      "Y",
			expectedFolioCreationDate: "09-08-2024",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockInput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_UNIT_BAL WHERE (.+)$").
					WillReturnRows(testCase.mockInput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_UNIT_BAL WHERE (.+)$").
					WillReturnError(errors.New("ORA Error"))
			}

			// Triggering Function
			actualOutput, err := suite.portfolioStore.GetUnblockedUnits(suite.ctx, 19, "IOGP", "12345", "12345678")

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedNoOfUnits, actualOutput.NoOfUnits)
				assert.Equal(t, testCase.expectedReinvestFlag, actualOutput.ReinvestFlag)
			}
		})
	}
}
