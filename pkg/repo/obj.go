package repo

import (
	"context"
	"kisaanSathi/pkg/logger"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Databases struct {
	PgDB *sqlx.DB
}

type DataObject struct {
	Databases Databases
	Cache     RedisInterface
}

func NewRepoObject(c context.Context) (DataObject, error) {
	logger.Log(c).Info("Creating new repository object")
	temp := DataObject{}
	var (
		readEbatestOracleDB *sqlx.DB
	)
	readEbatestOracleDB, err := PostgreSqlConnect()
	if err != nil {
		logger.Log(c).Error("Failed to get postgre connection", zap.Error(err))
		return temp, err
	}

	// If DB connection failed, log warning but continue with nil DB
	if readEbatestOracleDB == nil {
		logger.Log(c).Warn("PostgreSQL connection is nil - will use mock data for development/testing")
	}

	temp.Databases.PgDB = readEbatestOracleDB
	// redisObj, err := GetRedisObject(c)
	// if err != nil {
	// 	logger.Log(c).Error("Failed to get redis connection", zap.Error(err))
	// 	return temp, err
	// }
	// temp.Cache = redisObj

	return temp, nil
}
