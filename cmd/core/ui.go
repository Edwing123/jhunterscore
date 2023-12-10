package main

import (
	"io/fs"
	"log/slog"

	"edwingarcia.dev/github/jhunterscore/pkg/database/models"
	"edwingarcia.dev/github/jhunterscore/pkg/forms"
	"edwingarcia.dev/github/jhunterscore/ui"
	"github.com/gofiber/fiber/v2"
)

func NewViewsEngine() fiber.Views {
	// This is a workaround so that when rendering
	// views, we don't have to use the `html` prefix
	// in the view path.
	templatesFS, err := fs.Sub(ui.TemplateAssets, "html")
	if err != nil {
		slog.Error("Cannot read subdir", "err", err)
	}

	views := ui.NewEngine(templatesFS)
	return views
}

type ViewData struct {
	// The path of the page.
	Path string

	// The user data.
	User any

	// File related data.
	File  any
	Files []models.File

	// Companies related data.
	Company   any
	Companies []models.Company

	// Offers related data.
	Offer  any
	Offers []models.Offer

	// Locations related data
	Locations []models.Location

	// Navigation links.
	Links []Link

	// View errors.
	Errors forms.Errors
}

type Link struct {
	Name        string
	Path        string
	Icon        string
	IsAdminOnly bool
}

var links = []Link{
	{Name: "Inicio", Icon: "home", Path: "/admin"},
	{Name: "Ofertas", Icon: "newspaper", Path: "/admin/offers"},
	// {Name: "Recursos", Icon: "bagback", Path: "/admin/resources"},
	{Name: "Archivos", Icon: "files", Path: "/admin/files"},
	{Name: "Empresas", Icon: "buildings", Path: "/admin/companies", IsAdminOnly: true},
}
