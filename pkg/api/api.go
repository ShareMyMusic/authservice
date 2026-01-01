package api

import (
	"context"

	"github.com/sharemymusic/shared/pkg/env"
	"github.com/sharemymusic/shared/pkg/logger"
)

func Run(ctx context.Context) error {
	ctx = logger.CreateContext(ctx, env.Development)
	return nil
}
