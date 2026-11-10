package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"lorekeeper/api"
	"lorekeeper/api/routes"
	"lorekeeper/pkg/logging"
	"lorekeeper/pkg/validation"
	"os"
)

// @title Lorekeeper
// @version 0.1.1
// @description An AI knowledge assistant designed to help Dungeons & Dragons players organise, search and rediscover information from their campaign notes

// @contact.name Jared Brokenshire
// @contact.email jbrokenshire0306@gmail.com

// @BasePath /
func main() {
	// Load .env when present (ignored in Docker where env vars are injected directly).
	_ = godotenv.Load()

	server := api.NewServer()
	server.Echo.Validator = validation.NewCustomValidator(validator.New())
	server.Echo.HideBanner = true

	routes.RegisterRoutes(server)

	// Start Server
	logging.Logger.LogInfo().Msg(server.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))).Error())
}
