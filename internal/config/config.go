package config

import "github.com/sharemymusic/shared/pkg/env"

type Config struct {
	Env env.Env `env:"ENV" env-required:"true"`
}
