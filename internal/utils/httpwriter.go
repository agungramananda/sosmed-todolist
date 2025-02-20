package utils

import (
	"github.com/agungramananda/sosmed-todolist/internal/common/httpres"
	"github.com/labstack/echo/v4"
)

func WriteResponse(c echo.Context, code int, data any, msg string) error{
	return c.JSON(code, httpres.BaseResponse{
		Data: data,
		Message: msg,
	})
}