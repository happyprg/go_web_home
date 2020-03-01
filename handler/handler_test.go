package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{}
	if assert.NoError(t, h.Health(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, HealthOkMsg, rec.Body.String())
	}
}
