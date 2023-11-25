package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (core *Core) Setup() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:       "Nameless",
		ServerHeader:  "Go+FiberV2",
		CaseSensitive: true,
	})

	// Define global middlewares.
	app.Use(
		recover.New(),
		logger.New(),
		core.ManageSession,
	)

	api := app.Group("/api/v1")

	offers := api.Group("/offers")
	offers.Get("/", core.HandldeOffers)
	offers.Get("/:id<int>", core.HandldeOfferById)

	resources := api.Group("/resources")
	resources.Get("/", core.HandldeResources)
	resources.Get("/:id<int>", core.HandldeResourceById)

	return app
}
