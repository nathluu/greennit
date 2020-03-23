package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nathluu/greennit/internal/pkg/validator"
	"net/http"
)

type (
	Config struct {
		DefaultMiddlewares []echo.MiddlewareFunc
		Routes             []Route
	}

	Route struct {
		Desc        string
		Path        string
		Method      string
		Handler     echo.HandlerFunc
		Middlewares []echo.MiddlewareFunc
	}
)

func New(conf *Config) (http.Handler, error) {
	r := echo.New()
	r.Validator = validator.NewValidator()
	r.Use(conf.DefaultMiddlewares...)
	for _, route := range conf.Routes {
		r.Add(route.Method, route.Path, route.Handler, route.Middlewares...)
	}
	return r, nil
}
