package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nathluu/greennit/internal/pkg/config"
	"github.com/nathluu/greennit/internal/pkg/http/router"
	"net/http"
)

func NewRouter(config *config.Config) (http.Handler, error) {
	//authHandler := auth.NewHandler()
	routes := []router.Route{
		{
			Path:    "/api/v1",
			Method:  http.MethodGet,
			Handler: nil,
		},
	}
	//routes = append(routes, auth.)
	conf := router.Config{
		DefaultMiddlewares: getDefaultMiddlewares(),
		Routes:             routes,
	}

	r, err :=  router.New(&conf)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func getDefaultMiddlewares() []echo.MiddlewareFunc {
	mw := []echo.MiddlewareFunc{}
	lm := middleware.LoggerWithConfig(middleware.DefaultLoggerConfig)
	mw = append(mw, lm)

	bm := middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{
		Limit: "2M",
	})
	mw = append(mw, bm)
	return mw
}
