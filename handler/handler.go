package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const HealthOkMsg = "I'm alive"

type Handler struct {
}

func (h *Handler) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "serverTime", fmt.Sprintf("firstPR with %s", time.Now()))
}

func (h *Handler) Health(c echo.Context) error {
	return c.String(http.StatusOK, HealthOkMsg)
}
