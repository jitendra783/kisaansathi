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

type SchemeFlagsValidationSuite struct {
	suite.Suite
	ctx         context.Context
	sqlDB       *sql.DB
	gormDB      *gorm.DB
	sqlMock     sqlmock.Sqlmock
	schemeStore SchemeStore
}

func TestSchemeFlagsValidationSuite(t *testing.T) {
	suite.Run(t, new(SchemeFlagsValidationSuite))
}

func (suite *SchemeFlagsValidationSuite) SetupSuite() {
	logger.LoggerInit("", -1)

	suite.ctx = context.TODO()
	suite.sqlDB, suite.gormDB, suite.sqlMock = utils.NewMockDB()
	suite.schemeStore = NewSchemeStore(suite.gormDB)
}

func (suite *SchemeFlagsValidationSuite) TestValidateSchemeFlagsForPurchase() {
	testCases := []struct {
		desc             string
		isOfflineRequest bool
		mockOutput       *sqlmock.Rows
		expectedError    string
	}{
		{
			desc:             "SQLError",
			isOfflineRequest: true,
			mockOutput:       nil,
			expectedError:    "ORA Error",
		},
		{
			desc:             "OfflineNotEnabled",
			isOfflineRequest: true,
			mockOutput: sqlmock.NewRows([]string{"OfflineFlag", "OnlineFlag", "PurchaseAllowedFlag", "PurchaseFlag"}).
				AddRow("N", "Y", "N", "Y"),
			expectedError: "Scheme not offline enabled",
		}, {
			desc:             "OnlineNotEnabled",
			isOfflineRequest: false,
			mockOutput: sqlmock.NewRows([]string{"OfflineFlag", "OnlineFlag", "PurchaseAllowedFlag", "PurchaseFlag"}).
				AddRow("N", "N", "N", "Y"),
			expectedError: "Scheme not online enabled",
		}, {
			desc:             "PurchaseNotEnabled",
			isOfflineRequest: false,
			mockOutput: sqlmock.NewRows([]string{"OfflineFlag", "OnlineFlag", "PurchaseAllowedFlag", "PurchaseFlag"}).
				AddRow("N", "Y", "Y", "N"),
			expectedError: "This scheme is not enabled for Purchase",
		}, {
			desc:             "PurchaseNotAllowed",
			isOfflineRequest: false,
			mockOutput: sqlmock.NewRows([]string{"OfflineFlag", "OnlineFlag", "PurchaseAllowedFlag", "PurchaseFlag"}).
				AddRow("N", "Y", "N", "Y"),
			expectedError: "The scheme is closed for further transaction",
		}, {
			desc:             "PurchaseAllowed",
			isOfflineRequest: false,
			mockOutput: sqlmock.NewRows([]string{"OfflineFlag", "OnlineFlag", "PurchaseAllowedFlag", "PurchaseFlag"}).
				AddRow("N", "Y", "Y", "Y"),
			expectedError: "",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockOutput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnRows(testCase.mockOutput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnError(errors.New(testCase.expectedError))
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.ValidateSchemeFlagsForPurchase(suite.ctx, 19, "10GP", testCase.isOfflineRequest)

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput.PurchaseFlag.String(), "Y")
				assert.Equal(t, actualOutput.PurchaseAllowedFlag.String(), "Y")
				assert.Equal(t, actualOutput.OnlineFlag.String(), "Y")
			}
		})
	}
}

func (suite *SchemeFlagsValidationSuite) TestValidateSchemeFlagsForSIP() {
	testCases := []struct {
		desc             string
		isOfflineRequest bool
		mockOutput       *sqlmock.Rows
		expectedError    string
	}{
		{
			desc:             "SQLError",
			isOfflineRequest: true,
			mockOutput:       nil,
			expectedError:    "ORA Error",
		},
		{
			desc:             "OfflineNotEnabled",
			isOfflineRequest: true,
			mockOutput: sqlmock.NewRows([]string{"OfflineFlag", "OnlineFlag", "PurchaseAllowedFlag", "SIPFlag"}).
				AddRow("N", "Y", "N", "Y"),
			expectedError: "Scheme not offline enabled",
		}, {
			desc:             "OnlineNotEnabled",
			isOfflineRequest: false,
			mockOutput: sqlmock.NewRows([]string{"OfflineFlag", "OnlineFlag", "PurchaseAllowedFlag", "SIPFlag"}).
				AddRow("N", "N", "N", "Y"),
			expectedError: "Scheme not online enabled",
		}, {
			desc:             "SIPNotEnabled",
			isOfflineRequest: false,
			mockOutput: sqlmock.NewRows([]string{"OfflineFlag", "OnlineFlag", "PurchaseAllowedFlag", "SIPFlag"}).
				AddRow("N", "Y", "Y", "N"),
			expectedError: "This scheme is not enabled for SIP",
		}, {
			desc:             "PurchaseNotAllowed",
			isOfflineRequest: false,
			mockOutput: sqlmock.NewRows([]string{"OfflineFlag", "OnlineFlag", "PurchaseAllowedFlag", "SIPFlag"}).
				AddRow("N", "Y", "N", "Y"),
			expectedError: "The scheme is closed for further transaction",
		}, {
			desc:             "PurchaseAllowed",
			isOfflineRequest: false,
			mockOutput: sqlmock.NewRows([]string{"OfflineFlag", "OnlineFlag", "PurchaseAllowedFlag", "SIPFlag"}).
				AddRow("N", "Y", "Y", "Y"),
			expectedError: "",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockOutput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnRows(testCase.mockOutput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnError(errors.New(testCase.expectedError))
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.ValidateSchemeFlagsForSIP(suite.ctx, 19, "10GP", testCase.isOfflineRequest)

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput.SIPFlag.String(), "Y")
				assert.Equal(t, actualOutput.PurchaseAllowedFlag.String(), "Y")
				assert.Equal(t, actualOutput.OnlineFlag.String(), "Y")
			}
		})
	}
}

func (suite *SchemeFlagsValidationSuite) TestValidateSchemeFlagsForRedeem() {
	testCases := []struct {
		desc              string
		isSpecialInterval bool
		mockOutput        *sqlmock.Rows
		expectedError     string
	}{
		{
			desc:              "SQLError",
			isSpecialInterval: true,
			mockOutput:        nil,
			expectedError:     "ORA Error",
		},
		{
			desc:              "CloseEndedScheme",
			isSpecialInterval: true,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "RedeemFlag", "RedeemAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Close Ended", "N", "Y", "Y"),
			expectedError: "This scheme is close ended scheme and currently not available for redemption",
		}, {
			desc:              "ActivityNotAllowed",
			isSpecialInterval: false,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "RedeemFlag", "RedeemAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "N", "Y", "Y"),
			expectedError: "Requested activity not allowed on this scheme",
		}, {
			desc:              "NoFurtherTransaction",
			isSpecialInterval: false,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "RedeemFlag", "RedeemAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "N", "Y"),
			expectedError: "The scheme is closed for further transaction",
		}, {
			desc:              "NotSpecialInterval",
			isSpecialInterval: true,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "RedeemFlag", "RedeemAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "Y", "N"),
			expectedError: "Scheme is not Special Interval",
		}, {
			desc:              "RedeemAllowed",
			isSpecialInterval: true,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "RedeemFlag", "RedeemAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "Y", "Y"),
			expectedError: "",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockOutput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnRows(testCase.mockOutput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnError(errors.New(testCase.expectedError))
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.ValidateSchemeFlagsForRedeem(suite.ctx, 19, "10GP", testCase.isSpecialInterval)

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput.RedeemAllowedFlag.String(), "Y")
				assert.Equal(t, actualOutput.RedeemFlag.String(), "Y")
				assert.Equal(t, actualOutput.SpecialIntervalFlag.String(), "Y")
			}
		})
	}
}

func (suite *SchemeFlagsValidationSuite) TestValidateSchemeFlagsForSwitch() {
	testCases := []struct {
		desc              string
		isSpecialInterval bool
		mockOutput        *sqlmock.Rows
		expectedError     string
	}{
		{
			desc:              "SQLError",
			isSpecialInterval: true,
			mockOutput:        nil,
			expectedError:     "ORA Error",
		},
		{
			desc:              "CloseEndedScheme",
			isSpecialInterval: true,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "SwitchFlag", "SwitchAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Close Ended", "N", "Y", "Y"),
			expectedError: "This scheme is close ended scheme and currently not available for switch out",
		}, {
			desc:              "ActivityNotAllowed",
			isSpecialInterval: false,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "SwitchFlag", "SwitchAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "N", "Y", "Y"),
			expectedError: "Requested activity not allowed on this scheme",
		}, {
			desc:              "NoFurtherTransaction",
			isSpecialInterval: false,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "SwitchFlag", "SwitchAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "N", "Y"),
			expectedError: "The scheme is closed for further transaction",
		}, {
			desc:              "NotSpecialInterval",
			isSpecialInterval: true,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "SwitchFlag", "SwitchAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "Y", "N"),
			expectedError: "Scheme is not Special Interval",
		}, {
			desc:              "SwitchAllowed",
			isSpecialInterval: true,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "SwitchFlag", "SwitchAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "Y", "Y"),
			expectedError: "",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockOutput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnRows(testCase.mockOutput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnError(errors.New(testCase.expectedError))
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.ValidateSchemeFlagsForSwitch(suite.ctx, 19, "10GP", testCase.isSpecialInterval)

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput.SwitchAllowedFlag.String(), "Y")
				assert.Equal(t, actualOutput.SwitchFlag.String(), "Y")
				assert.Equal(t, actualOutput.SpecialIntervalFlag.String(), "Y")
			}
		})
	}
}

func (suite *SchemeFlagsValidationSuite) TestValidateSchemeFlagsForSTP() {
	testCases := []struct {
		desc              string
		isSpecialInterval bool
		isBoosterSTP      bool
		isDematHolding    bool
		mockOutput        *sqlmock.Rows
		mockRegistrar     *sqlmock.Rows
		expectedError     string
	}{
		{
			desc:              "SQLError",
			isSpecialInterval: true,
			mockOutput:        nil,
			expectedError:     "ORA Error",
		}, {
			desc:         "CloseEndedSchemeBoosterSTP",
			isBoosterSTP: true,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "BoosterSTPFlag", "STPOutFlag", "SwitchAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Close Ended", "N", "Y", "Y", "Y"),
			expectedError: "This scheme is close ended scheme and currently not available for STP out",
		}, {
			desc:         "CloseEndedScheme",
			isBoosterSTP: false,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "BoosterSTPFlag", "STPOutFlag", "SwitchAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Close Ended", "Y", "N", "Y", "Y"),
			expectedError: "This scheme is close ended scheme and currently not available for STP out",
		}, {
			desc:         "BoosterSTPNotEnabled",
			isBoosterSTP: true,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "BoosterSTPFlag", "STPOutFlag", "SwitchAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "N", "Y", "Y", "Y"),
			expectedError: "This Scheme is not enabled for time the market STP",
		}, {
			desc:         "STPNotEnabled",
			isBoosterSTP: false,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "BoosterSTPFlag", "STPOutFlag", "SwitchAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "N", "Y", "Y"),
			expectedError: "Scheme not enabled for STP Out",
		}, {
			desc:              "NoFurtherTransaction",
			isSpecialInterval: false,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "BoosterSTPFlag", "STPOutFlag", "SwitchAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "Y", "N", "Y"),
			expectedError: "The scheme is closed for further transaction",
		}, {
			desc:              "NotSpecialInterval",
			isSpecialInterval: true,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "BoosterSTPFlag", "STPOutFlag", "SwitchAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "Y", "Y", "N"),
			expectedError: "Scheme is not Special Interval",
		}, {
			desc:              "InvalidRegistrar",
			isSpecialInterval: true,
			isDematHolding:    true,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "BoosterSTPFlag", "STPOutFlag", "SwitchAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "Y", "Y", "Y"),
			mockRegistrar: sqlmock.NewRows([]string{"registrar"}).AddRow("KARVY"),
			expectedError: "STP order not allowed",
		}, {
			desc:              "STPAllowed",
			isSpecialInterval: true,
			isDematHolding:    true,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "BoosterSTPFlag", "STPOutFlag", "SwitchAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "Y", "Y", "Y"),
			mockRegistrar: sqlmock.NewRows([]string{"registrar"}).AddRow("KAR"),
			expectedError: "",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockOutput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnRows(testCase.mockOutput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnError(errors.New(testCase.expectedError))
			}

			if testCase.mockRegistrar != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_COMPANIES WHERE (.+)$").
					WillReturnRows(testCase.mockRegistrar)
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.ValidateSchemeFlagsForSTP(suite.ctx, 19, "10GP", testCase.isSpecialInterval, testCase.isBoosterSTP, testCase.isDematHolding)

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput.SwitchAllowedFlag.String(), "Y")
				assert.Equal(t, actualOutput.BoosterSTPFlag.String(), "Y")
				assert.Equal(t, actualOutput.STPOutFlag.String(), "Y")
				assert.Equal(t, actualOutput.SpecialIntervalFlag.String(), "Y")
			}
		})
	}
}

func (suite *SchemeFlagsValidationSuite) TestValidateSchemeFlagsForSWP() {
	testCases := []struct {
		desc            string
		specialInterval bool
		mockOutput      *sqlmock.Rows
		expectedError   string
	}{
		{
			desc:            "SQLError",
			specialInterval: true,
			mockOutput:      nil,
			expectedError:   "ORA Error",
		}, {
			desc:            "ActivityNotAllowed",
			specialInterval: false,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "SWPFlag", "RedeemAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "N", "Y", "Y"),
			expectedError: "Requested activity not allowed on this scheme",
		}, {
			desc:            "NoFurtherTransaction",
			specialInterval: false,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "SWPFlag", "RedeemAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "N", "Y"),
			expectedError: "The scheme is closed for further transaction",
		}, {
			desc:            "NotSpecialInterval",
			specialInterval: true,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "SWPFlag", "RedeemAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "Y", "N"),
			expectedError: "Scheme is not Special Interval",
		}, {
			desc:            "SWPAllowed",
			specialInterval: true,
			mockOutput: sqlmock.NewRows([]string{"CloseFlag", "SWPFlag", "RedeemAllowedFlag", "SpecialIntervalFlag"}).
				AddRow("Open Ended", "Y", "Y", "Y"),
			expectedError: "",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.desc, func(t *testing.T) {
			// Mocking and Setting Expected Result
			if testCase.mockOutput != nil {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnRows(testCase.mockOutput)
			} else {
				suite.sqlMock.
					ExpectQuery("^SELECT (.+) FROM MF_SCHEME_MASTER, MF_RECOMMENDATIONS WHERE (.+)$").
					WillReturnError(errors.New(testCase.expectedError))
			}

			// Triggering Function
			actualOutput, err := suite.schemeStore.ValidateSchemeFlagsForSWP(suite.ctx, 19, "10GP", testCase.specialInterval)

			// Validations
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, actualOutput.RedeemAllowedFlag.String(), "Y")
				assert.Equal(t, actualOutput.SWPFlag.String(), "Y")
				assert.Equal(t, actualOutput.SpecialIntervalFlag.String(), "Y")
			}
		})
	}
}
