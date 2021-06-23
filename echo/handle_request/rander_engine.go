package main

import (
	"embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

//go:embed templates
var content embed.FS

type RenderEngine struct {
	tmpl *template.Template
}

func (e *RenderEngine) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return e.tmpl.ExecuteTemplate(w, name, data)
}

func newRenderEngine() *RenderEngine {
	return &RenderEngine{
		tmpl: template.Must(template.ParseFS(content, "templates/*.html")),
	}
}
