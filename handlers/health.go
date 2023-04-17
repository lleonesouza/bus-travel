package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

func (u *HealthHandler) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, `{"status":"ok"}`)
}
