package main

import (
	"context"
	"fmt"
	"kisaanSathi/api"
	"kisaanSathi/pkg/config"
	"kisaanSathi/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var environment string
	host := os.Getenv("SERVER_HOST")
	if host != "" {
		environment = "server"
	} else {
		// Check if a custom environment file is provided
		if len(os.Args) == 2 {
			environment = os.Args[1] // developer custom file
		} else {
			environment = "local" // default to local environment
		}
		config.Load(environment)
		if err := api.Start(); err != nil {
			log.Fatal("Failed to start server, err:", err)
			os.Exit(1)
		}

		addShutdownHook()
	}
}

// addShutdownHook sets up a signal handler to gracefully shut down the server
func addShutdownHook() {

	// Listen for system signals to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// Block until a signal is received
	<-quit
	logger.Log().Info("Quit/Interrupt signal detected. Gracefully closing connections")

	// Shutdown the server
	api.ShutdownRouter()
	api.CloseDatabase()

	ctx := context.Background()

	logger.Log(ctx).Info(fmt.Sprintf("All done! Wrapping up here for PID: %d", os.Getpid()))
}
