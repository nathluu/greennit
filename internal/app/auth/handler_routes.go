package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/nathluu/greennit/internal/pkg/http"
)

func (h *Handler) Routes() []http.Route {
	return []http.Route{
		{
			Path:    "/api/v1/auth",
			Method:  echo.POST,
			Handler: h.Auth,
		},
	}
}
