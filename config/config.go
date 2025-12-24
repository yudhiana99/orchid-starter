package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"path/filepath"
)

func GetLocalConfig() *LocalConfig {

	// Load .env file for local development
	_ = godotenv.Load(filepath.Join("..", "..", ".env"))

	var localConfig LocalConfig
	if err := env.Parse(&localConfig); err != nil {
		panic(err)
	}
	return &localConfig
}
