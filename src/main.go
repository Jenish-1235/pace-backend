package main

import (
	"context"
	utils "pace-backend/src/utils"
	"pace-backend/src/utils/postgres"
)

func main() {
	ctx := context.Background()

	cfg := utils.GetConfig()

	utils.InitLogger(cfg.LogLevel, cfg.Environment)
	logger := utils.GetLogger()
	
	logger.Infof("Application %s starting on port %d in %s mode", cfg.Name, cfg.HttpPort, cfg.Environment)

	_, err := postgres.Init(ctx, postgres.Config{
		DatabaseURI: cfg.PostgresDBUri,
		MaxConns:    cfg.MaxConns,
		MinConns:    cfg.MinConns,
	})
	if err != nil {
		logger.Fatalf("failed to initialize postgres: %v", err)
	}
	defer postgres.Close()
	
	logger.Info("Postgres connection pool initialized successfully")

	err = postgres.RunMigrations(ctx, "src/migrations")
	if err != nil {
	  logger.Fatalf("database migrations failed: %v", err)
	}
	logger.Info("Database migrations applied successfully")
	
	ctx = utils.LoggerWithContext(ctx, logger)
}
