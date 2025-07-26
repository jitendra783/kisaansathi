package api

import (
	"context"
	"kisaanSathi/pkg/config"
	"kisaanSathi/pkg/logger"
	serv "kisaanSathi/pkg/services"

	"fmt"
	"kisaanSathi/pkg/repo"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

var srv *http.Server
var ctx context.Context
var databases []*gorm.DB

// starts the server with initializations
//
//	initializes logs
//	creates global context
//	connects databases
//	connects redis
//	creates versioned service objects
func Start() error {
	ctx = context.Background()

	config := config.GetConfig()
	logLevel, err := strconv.Atoi(config.GetString("log.Level"))
	if err != nil {
		log.Fatal("Invalid log config: ", err)
	}
	logger.LoggerInit(config.GetString("log.path"), zapcore.Level(logLevel))
	
	repoObj, err := repo.NewRepoObject(ctx)
	if err != nil {
		logger.Log().Error("Failed to create repo object", zap.Error(err))
		return err
	}
	serviceObj := serv.NewServiceObject(repoObj)
	startRouter(serviceObj)
	return nil
}


func startRouter(obj serv.ServiceLayer) {
	srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetConfig().GetInt("server.port")),
		Handler: getRouter(obj, logger.Log()), //getRouter set the api specs for version-1 routes
	}
	// run api router
	logger.Log().Info("starting router")
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log().Fatal("Error starting server", zap.Error(err))
		}
	}()
}

// stops the router running in the go routine.
//
//	uses Shutdown() function of native http server library
//	internally defaults a 5 seconds context timeout.
//		timeoutCtx,_ := context.WithTimeout(ctx, 5*time.Second)
//		srv.Shutdown(timeoutCtx)
func ShutdownRouter() {
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	logger.Log().Info("Shutting down router START")
	defer logger.Log().Info("Shutting down router END")
	if err := srv.Shutdown(timeoutCtx); err != nil {
		logger.Log().Fatal("Server forced to shutdown", zap.Error(err))
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-timeoutCtx.Done():
		log.Println("timeout of 2 seconds.")
	}
}

// closes all database connections
//
//	closes each database connection made which are saved globally
//	logs error if unable to close
//		function used: *sql.DB.Close()
func CloseDatabase() {
	logger.Log().Info("disconnecting databases START")
	defer logger.Log().Info("disconnecting databases END")
	for _, database := range databases {
		db, _ := database.DB()
		if db != nil {
			err := db.Close()
			if err != nil {
				logger.Log().Error("unable to close db", zap.Error(err))
			}
		} else {
			logger.Log().Error("unable to close db as connection is nil")

		}
	}
}
