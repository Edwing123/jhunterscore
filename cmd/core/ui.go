package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func NewViewsEngine() fiber.Views {
	views := html.New(
		"./ui/html",
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
