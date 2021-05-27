package main

import (
	"flag"

	"github.com/alseiitov/bookstore/internal/app"
	"github.com/alseiitov/dotenv"
)

func main() {
	configPath := flag.String("config-path", "./configs/config.json", "Path to the config file")
	flag.Parse()

	dotenv.Load()

	app.Run(*configPath)
}
