package router

import (
	"rest-gateway/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitilizeAuthRoutes(app *fiber.App) {
	authHandler := handlers.AuthHandler{}
	api := app.Group("/auth", logger.New())
	api.Post("/:stationName/consume", handlers.ConsumeHandleMessage())
	api.Post("/:stationName/consume/ack", handlers.AcknowledgeMessage())
	api.Post("/authenticate", authHandler.Authenticate)
	api.Post("/refreshToken", authHandler.RefreshToken)
}
