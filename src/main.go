package main

import (
	"context"
	utils "pace-backend/src/utils"
)

func main() {
	ctx := context.Background()

	cfg := utils.GetConfig()

	utils.InitLogger(cfg.LogLevel, cfg.Environment)
	logger := utils.GetLogger()
	
	logger.Infof("Application %s starting on port %d in %s mode", cfg.Name, cfg.HttpPort, cfg.Environment)
	
	ctx = utils.LoggerWithContext(ctx, logger)
}
