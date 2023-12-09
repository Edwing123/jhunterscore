package ui

import "embed"

//go:embed static
var StaticAssets embed.FS

//go:embed html
var TemplateAssets embed.FS
