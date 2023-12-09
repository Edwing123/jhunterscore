package main

import (
	"io/fs"
	"log/slog"
	"net/http"

	"edwingarcia.dev/github/jhunterscore/ui"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func NewViewsEngine() fiber.Views {
	// This is a workaround so that when rendering
	// views, we don't have to use the `html` prefix
	// in the view path.
	templatesFS, err := fs.Sub(ui.TemplateAssets, "html")
	if err != nil {
		slog.Error("Cannot read subdir", "err", err)
	}

	views := html.NewFileSystem(
		http.FS(templatesFS),
		".html",
	)

	return views
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
