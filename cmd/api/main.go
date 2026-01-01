package main

import (
	"context"

	"github.com/sharemymusic/authservice/pkg/api"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	err := api.Run(ctx)
	if err != nil {
		zap.L().Fatal("failed to run api", zap.Error(err))
	}
}
