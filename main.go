package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func main() {
	app := fiber.New()

	app.Use(loggingMiddleware)

	app.Get("/", helloWorldHandler)

	if err := app.Listen(":8080"); err != nil {
		logger.Fatal("error while starting the server", zap.Error(err))
	}
}

func helloWorldHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello World!")
}
func loggingMiddleware(ctx *fiber.Ctx) error {
	logger.Info("An api call made.", zap.String("path", ctx.Path()))
	return ctx.Next()
}
