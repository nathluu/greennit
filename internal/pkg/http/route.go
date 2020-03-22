package http

import "github.com/labstack/echo/v4"

type (
	Route struct {
		Desc        string
		Path        string
		Method      string
		Handler     echo.HandlerFunc
		Middlewares []echo.MiddlewareFunc
	}
)
