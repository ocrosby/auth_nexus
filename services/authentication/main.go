package main

import (
	"github.com/ocrosby/auth-nexus/internal/authentication"
	"github.com/ocrosby/auth-nexus/internal/authentication/router"
	"github.com/ocrosby/auth-nexus/pkg/logger"
)

func main() {
	var (
		err error
	)

	mux := router.SetupRoutes()
	logger := logger.NewLogger()
	service := authentication.NewServer(logger, mux)

	// Log an info message
	logger.Info("Application started", map[string]interface{}{
		"version": "1.0.0",
	})

	// Log an error message
	logger.Error("An error occurred", map[string]interface{}{
		"error": "sample error message",
	})

	// Log a debug message
	logger.Debug("Debugging application", map[string]interface{}{
		"debug": "debug details",
	})

	if err = service.Run(); err != nil {
		logger.Error("An error occurred", map[string]interface{}{
			"error": err.Error(),
		})
	}
}
