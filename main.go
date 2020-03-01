package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! 2020-03-01 12:52")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
