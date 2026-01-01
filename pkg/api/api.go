package api

import (
	"context"
	"fmt"

	"github.com/sharemymusic/authservice/internal/config"
	"github.com/sharemymusic/shared/pkg/configloader"
	"github.com/sharemymusic/shared/pkg/env"
	"github.com/sharemymusic/shared/pkg/logger"
	"go.uber.org/zap"
)

func Run(ctx context.Context) error {
	var cfg config.Config

	err := configloader.Load(&cfg, "config/.env")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	ctx = logger.CreateContext(ctx, env.Development)

	log := logger.FromContext(ctx)
	log.Debug("starting api", zap.String("env", cfg.Env.String()))

	// Main logic

	log.Info("shutting down api", zap.String("env", cfg.Env.String()))

	defer func() {
		_ = log.Sync()
	}()

	return nil
}
