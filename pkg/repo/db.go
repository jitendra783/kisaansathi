package repo

import (
	"context"
	"kisaanSathi/pkg/config"
	elog "kisaanSathi/pkg/logger"
	e "kisaanSathi/pkg/network"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	// "gorm.io/gorm/schema"

	"time"
)

type dbLogger struct{}

func PostgreSqlConnect() (*gorm.DB, error) {
	elog.Log().Info("Connecting to PostgreSQL database")

	c := config.GetConfig()

	dsn := "host=" + c.GetString("database.host") +
		" user=" + c.GetString("database.user") +
		" password=" + c.GetString("database.password") +
		" dbname=" + c.GetString("database.db") +
		" port=" + c.GetString("database.port") 
		
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: customLogger(),
	})
	if err != nil {
		elog.Log().Error("failed to connect postgreSQL connection", zap.Error(err), zap.String("connStr", dsn))
		return nil, e.ApiErrors.PostgresDBConnError

	}
	elog.Log().Info("postgre Database Connected")
	return db, nil
}

func customLogger() logger.Interface {
	return dbLogger{}
}

func (d dbLogger) Error(ctx context.Context, data string, others ...interface{}) {
	elog.Log(ctx).Info("database", zap.String("error", data), zap.Any("description", others))
}

func (d dbLogger) Info(ctx context.Context, data string, others ...interface{}) {
	elog.Log(ctx).Info("database", zap.String("msg", data), zap.Any("description", others))
}

func (d dbLogger) Warn(ctx context.Context, data string, others ...interface{}) {
	elog.Log(ctx).Info("database", zap.String("msg", data), zap.Any("description", others))
}

func (d dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	query, others := fc()
	if err != nil {
		elog.Log(ctx).Info("database", zap.String("query", query), zap.Any("rows-affected", others), zap.Error(err))
	} else {
		elog.Log(ctx).Info("database", zap.String("query", query), zap.Any("rows-affected", others))
	}
}

func (d dbLogger) LogMode(l logger.LogLevel) logger.Interface {
	return d
}
