package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "I'm alive")
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! add static")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
