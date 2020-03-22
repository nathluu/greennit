package api

import (
	"github.com/labstack/echo/v4"
)

func NewRouter() (*echo.Echo, error) {
	e := echo.New()
	return e, nil
}
