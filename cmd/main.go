package main

import (
	"simple_cryptographer/app"
	"simple_cryptographer/config"
)

func main() {
	cfg := config.New()
	app.Run(cfg)
}
