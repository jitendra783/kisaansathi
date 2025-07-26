package repo

import (
	"context"
	"kisaanSathi/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Databases struct {
	PgDB *gorm.DB
}

type DataObject struct {
	Databases Databases
	Cache     RedisInterface
}

func NewRepoObject(c context.Context) (DataObject, error) {
	logger.Log(c).Info("Creating new repository object")
	temp := DataObject{}
    var (
		readEbatestOracleDB *gorm.DB
	)
	// readEbatestOracleDB, err := PostgreSqlConnect()
	// if err != nil {
	// 	logger.Log(c).Error("Failed to get oracle connection", zap.Error(err))
	// 	return temp, err
	// }
	temp.Databases.PgDB = readEbatestOracleDB
	redisObj, err := GetRedisObject(c)
	if err != nil {
		logger.Log(c).Error("Failed to get redis connection", zap.Error(err))
		return temp, err
	}
	temp.Cache = redisObj

	return temp, nil
}
