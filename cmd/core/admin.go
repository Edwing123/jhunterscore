package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func (core *Core) SetupAdmin(app *fiber.App) {
	admin := app.Group("/admin")

	admin.Use(
		helmet.New(),
	)

	admin.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("pages/admin/auth/login", nil)
	})

	admin.Use(
		core.RequireAuth,
	)

	admin.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/admin/home/index", nil)
	})

	// Offers pages.
	offers := admin.Group("/offers")

	offers.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/admin/offers/index", nil)
	})

	// Resources pages.
	resources := admin.Group("/resources")

	resources.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/admin/resources/index", nil)
	})

	// Files pages.
	files := admin.Group("/files")

	files.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/admin/files/index", nil)
	})

	// Companies pages.
	companies := admin.Group("/companies", core.RequireAdmin)

	companies.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/admin/companies/index", nil)
	})
}
