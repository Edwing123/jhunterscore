package main

import (
	"io/fs"
	"log/slog"

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
	Path  string
	User  any
	Links []Link
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
	{Name: "Recursos", Icon: "bagback", Path: "/admin/resources"},
	{Name: "Archivos", Icon: "files", Path: "/admin/files"},
	{Name: "Empresas", Icon: "buildings", Path: "/admin/companies", IsAdminOnly: true},
}
