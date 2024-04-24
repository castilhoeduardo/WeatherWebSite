package main

import (
	"io"
	"text/template"
	"wheater/cmd/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRender struct {
	templates *template.Template
  }
  
  func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
  }

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Renderer = &TemplateRender{
		templates: template.Must(template.ParseGlob("web/templates/*.html")),
	  }

	e.GET("/", handlers.GetCity) //TEMPLATE RENDER
	e.GET("/health", handlers.Health)

	e.Static("/static", "web/static")

	e.Logger.Fatal(e.Start(":8000"))
}