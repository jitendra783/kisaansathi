package db

import (
	"context"
	"database/sql"
	"errors"
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/utils"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type AnalyticsSuite struct {
	suite.Suite
	ctx            context.Context
	sqlDB          *sql.DB
	gormDB         *gorm.DB
	sqlMock        sqlmock.Sqlmock
	analyticsStore AnalyticsStore
}

func TestAnalyticsSuite(t *testing.T) {
	suite.Run(t, new(AnalyticsSuite))
}

func (suite *AnalyticsSuite) SetupSuite() {
	logger.LoggerInit("", -1)

	suite.ctx = context.TODO()
	suite.sqlDB, suite.gormDB, suite.sqlMock = utils.NewMockDB()
	suite.analyticsStore = NewAnalyticsStore(suite.gormDB)
}

func (suite *AnalyticsSuite) TearDownSuite() {
}

func (suite *AnalyticsSuite) SetupTest() {
}

func (suite *AnalyticsSuite) TearDownTest() {
}

func (suite *AnalyticsSuite) TestUpdateCampaignClickEvent_ReturnSQLError() {
	// Mocking and Setting Expected Result
	suite.sqlMock.ExpectBegin()
	suite.sqlMock.
		ExpectExec("^INSERT INTO MCM_MF_CMPGN_MTCH VALUES (.+)$").
		WillReturnError(errors.New("ORA Error"))
	suite.sqlMock.ExpectRollback()

	// Triggering Function
	err := suite.analyticsStore.UpdateCampaignClickEvent(suite.ctx, "000000000")

	// Validations
	suite.EqualError(err, "ORA Error")
}

func (suite *AnalyticsSuite) TestUpdateCampaignClickEvent_ReturnSuccess() {
	// Mocking and Setting Expected Result
	suite.sqlMock.ExpectBegin()
	suite.sqlMock.
		ExpectExec("^INSERT INTO MCM_MF_CMPGN_MTCH VALUES (.+)$").
		WillReturnResult(sqlmock.NewResult(0, 1))
	suite.sqlMock.ExpectCommit()

	// Triggering Function
	err := suite.analyticsStore.UpdateCampaignClickEvent(suite.ctx, "000000000")

	// Validations
	suite.NoError(err)
}
