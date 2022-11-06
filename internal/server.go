package internal

import (
	"github.com/request-dumper/rd/internal/logging"

	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

const appName = "Request-Dumper"

func StartServer(port int) {

	app := fiber.New(fiber.Config{
		AppName:               appName,
		Prefork:               false,
		DisableKeepalive:      true,
		DisableStartupMessage: true,
		Network:               fiber.NetworkTCP,
	})

	app.Use(logging.ZerologFiberMiddleware())
	app.Use(ok)

	log.Info().
		Int("port", port).
		Msgf(fmt.Sprintf("%s started", appName))
	log.Fatal().
		Err(app.Listen(fmt.Sprintf(":%d", port))).
		Msg("Application was stopped")
}

func ok(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
