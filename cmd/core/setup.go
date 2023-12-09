package main

import (
	"io/fs"
	"log/slog"
	"net/http"

	"edwingarcia.dev/github/jhunterscore/pkg/database/models"
	"edwingarcia.dev/github/jhunterscore/ui"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (core *Core) Setup() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:       "Nameless",
		ServerHeader:  "Go+FiberV2",
		CaseSensitive: true,
		Views:         NewViewsEngine(),
		ViewsLayout:   "layouts/base",
	})

	// Define global middlewares.
	app.Use(
		recover.New(),
		logger.New(),
		core.ManageSession,
	)

	// Path for static files.
	assetsFS, err := fs.Sub(ui.StaticAssets, "static")
	if err != nil {
		slog.Error("Cannot read subdir", "err", err)
	}

	app.Use("/static", filesystem.New(filesystem.Config{
		Root: http.FS(assetsFS),
	}))

	core.SetupAdmin(app)

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

func (core *Core) SetupAdmin(app *fiber.App) {
	admin := app.Group("/admin")

	// Protect UI pages by setting several HTTP headers.
	admin.Use(helmet.New())

	user := models.User{
		Role:     "admin",
		Username: "Edwing123",
	}

	admin.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/admin/home/index", fiber.Map{
			"Path":  c.Path(),
			"User":  user,
			"Links": links,
		})
	})

	admin.Get("/auth/login", func(c *fiber.Ctx) error {
		return c.Render("pages/admin/auth/login", nil, "")
	})
}
