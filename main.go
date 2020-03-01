package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "I'm alive")
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! 2020-03-01 13:53(KST) integrated with skaffold successfully, it's my fault")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
