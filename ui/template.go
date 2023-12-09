package ui

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
)

type Engine struct {
	templates   map[string]*template.Template
	templatesFS fs.FS
}

func (e *Engine) Load() error {
	templates, err := createTemplatesMap(e.templatesFS)
	if err != nil {
		return err
	}

	e.templates = templates
	return nil
}

func (e *Engine) Render(output io.Writer, name string, data any, layouts ...string) error {
	t, ok := e.templates[name]
	if !ok {
		return fmt.Errorf("template %s does not exist", name)
	}

	err := t.ExecuteTemplate(output, "layouts/base", data)
	return err
}

func NewEngine(templateFS fs.FS) *Engine {
	return &Engine{
		templatesFS: templateFS,
	}
}

// Create templates map from views found in "[dir]/views".
func createTemplatesMap(templatesFS fs.FS) (map[string]*template.Template, error) {
	// Get views file paths.
	views, err := fs.Glob(templatesFS, "pages/**/**/*")
	if err != nil {
		return nil, err
	}

	// Create map with a length of len(views).
	templates := make(map[string]*template.Template, len(views))

	// Iterable on the views files and construct a template set that inclues:
	// - the view
	// - components
	// - layouts
	for _, view := range views {
		ext := filepath.Ext(view)
		name := strings.TrimSuffix(view, ext)

		// Parse view file.
		ts, err := template.New(name).ParseFS(templatesFS, view)
		if err != nil {
			return nil, err

		}

		// Parse layouts files.
		ts, err = ts.ParseFS(templatesFS, "layouts/*.html")
		if err != nil {
			return nil, err

		}

		// Parse components files.
		ts, err = ts.ParseFS(templatesFS, "components/*.html")
		if err != nil {
			return nil, err

		}

		// Add template set to the map.
		templates[name] = ts
	}

	return templates, nil
}
