package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

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
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "serverTime", fmt.Sprintf("%s", time.Now()))
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "I'm alive")
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":1323"))
}
