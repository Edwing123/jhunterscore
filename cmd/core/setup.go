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

	api := app.Group("/api/")
	v1 := api.Group("/v1")

	offers := v1.Group("/offers")
	offers.Get("/", core.HandldeOffers)
	offers.Get("/:id<int>", core.HandldeOfferById)

	resources := v1.Group("/resources")
	resources.Get("/", core.HandldeResources)
	resources.Get("/:id<int>", core.HandldeResourceById)

	companies := v1.Group("/companies")
	companies.Get("/", core.HandldeCompanies)

	return app
}
