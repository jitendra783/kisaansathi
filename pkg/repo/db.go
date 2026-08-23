package repo

import (
	"context"
	"fmt"
	"kisaanSathi/pkg/config"
	elog "kisaanSathi/pkg/logger"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"

	"time"
)

type dbLogger struct{}

func PostgreSqlConnect() (*sqlx.DB, error) {
	elog.Log().Info("Connecting to PostgreSQL database")

	c := config.GetConfig()
	sslMode := c.GetString("database.sslmode")
	if sslMode == "" {
		sslMode = "disable"
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s channel_binding=%s search_path=%s",
		c.GetString("database.host"),
		c.GetString("database.port"),
		c.GetString("database.user"),
		c.GetString("database.password"),
		c.GetString("database.database"),
		sslMode,
		c.GetString("database.channel_binding"),
		c.GetString("database.schema"),
	)

	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		elog.Log().Error(
			"Failed to create PostgreSQL connection",
			zap.Error(err),
		)
		SetDBStatus(false, err.Error())
		return nil, err
	}

	// Connection Pool
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	if err := db.Ping(); err != nil {
		elog.Log().Error(
			"Failed to ping PostgreSQL",
			zap.Error(err),
		)
		SetDBStatus(false, err.Error())
		return nil, err
	}

	// Set PostgreSQL schema
	schema := c.GetString("database.schema")

	if schema != "" {
		_, err := db.Exec(
			fmt.Sprintf(`SET search_path TO "%s"`, schema),
		)
		if err != nil {
			elog.Log().Error(
				"Failed to set PostgreSQL schema",
				zap.Error(err),
			)
			SetDBStatus(false, err.Error())
			return nil, err
		}

		elog.Log().Info(
			"PostgreSQL schema set successfully",
			zap.String("schema", schema),
		)
	}

	elog.Log().Info("PostgreSQL Database Connected Successfully")

	SetDBStatus(true, "")

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
