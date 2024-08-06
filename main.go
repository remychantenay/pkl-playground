package main

import (
	"context"

	"github.com/remychantenay/pkl-playground/internal/config"
)

const appName = "Cortex"

var (
	// globalCfg holds the application config.
	globalCfg *config.Config
)

func main() {
	loadConfig()
}

func loadConfig() {
	cfg, err := config.LoadFromPath(context.Background(), "pkl/config.pkl")
	if err != nil {
		panic(err)
	}

	globalCfg = cfg
}
