package exceptions

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func CustomHTTPErrorHandler(logger zerolog.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var report *echo.HTTPError

		if httpErr, ok := err.(*echo.HTTPError); ok {
			report = httpErr
		} else if invErr, ok := err.(InvariantError); ok {
			report = echo.NewHTTPError(http.StatusBadRequest, invErr.Message)
		} else if notFoundErr, ok := err.(NotFoundError); ok {
			report = echo.NewHTTPError(http.StatusNotFound, notFoundErr.Message)
		} else if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range castedObject {
				var message string
				switch fieldErr.Tag() {
				case "required":
					message = fmt.Sprintf("%s is required", fieldErr.Field())
				case "min":
					message = fmt.Sprintf("%s value must be at least %s", fieldErr.Field(), fieldErr.Param())
				case "max":
					message = fmt.Sprintf("%s value must be at most %s", fieldErr.Field(), fieldErr.Param())
				default:
					message = fmt.Sprintf("%s is invalid", fieldErr.Field())
				}
				report = echo.NewHTTPError(http.StatusBadRequest, message)
			}
		} else {
			logger.Error().Err(err).Msg(err.Error())
			report = echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
		}

		logger.Error().
			Int("status", report.Code).
			Str("error", fmt.Sprintf("%v", report.Message)).
			Str("method", c.Request().Method).
			Str("uri", c.Request().RequestURI).
			Str("remote_ip", c.RealIP()).
			Msg("HTTP error occurred")

		if err := c.JSON(report.Code, report); err != nil {
			logger.Error().Err(err).Msg("Failed to send JSON response")
		}
	}
}
