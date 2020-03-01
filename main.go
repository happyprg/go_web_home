package main

import (
	"html/template"
	"io"

	"github.com/happyprg/go_web_home/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {

	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("template/views/*.html")),
	}
	e.Renderer = t
	h := handler.Handler{}
	e.GET("/", h.Index)
	e.GET("/health", h.Health)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":1323"))
}
