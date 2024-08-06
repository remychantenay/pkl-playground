package main

import (
	"context"
	"fmt"

	"github.com/remychantenay/pkl-playground/internal/config"
)

func main() {
	loadConfig()
}

func loadConfig() {
	cfg, err := config.LoadFromPath(context.Background(), "pkl/config.pkl")
	if err != nil {
		panic(err)
	}

	println(fmt.Sprintf("SomeInteger: %d", cfg.SomeInteger))
	println(fmt.Sprintf("SomeString: %s", cfg.SomeString))
	println(fmt.Sprintf("Tabs (list): %v", cfg.Tabs))
	println(fmt.Sprintf("Components (map): %v", cfg.Components))
	println(fmt.Sprintf("FullList (list): %v", cfg.FullList))
	println(fmt.Sprintf("Multiplied (list): %v", cfg.MultipliedList))
	println(fmt.Sprintf("Filtered (list): %v", cfg.FilteredList))
}
