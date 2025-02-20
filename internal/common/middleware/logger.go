package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func RequestLogger(logger *zerolog.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Int("status", v.Status).
				Str("URI", v.URI).
				Str("method", v.Method).
				Str("remote_ip", v.RemoteIP).
				Msg("request")
			return nil
		},
	})
}