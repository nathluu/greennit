package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/nathluu/greennit/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/auth",
			Method:  echo.POST,
			Handler: h.Auth,
		},
	}
}
