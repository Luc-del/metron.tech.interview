package api

import (
	"github.com/gofiber/fiber/v2"
)

var (
	statusRoute   = "/status"
	costRoute     = "/cost"
	discountRoute = "/discountSettings"
)

func BuildAPI(version string) *fiber.App {
	app := fiber.New(
		fiber.Config{
			Immutable: true,
		},
	)

	app.Get("/", status(version))
	app.Get(statusRoute, status(version))

	app.Post(costRoute, cost())
	app.Post(discountRoute, setDiscount())

	return app
}
