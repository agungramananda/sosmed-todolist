package platforms

import (
	"context"
	"net/http"

	"github.com/agungramananda/sosmed-todolist/internal/utils"
	"github.com/labstack/echo/v4"
)

type GetAllPlatformsHandler func (context.Context, *PlatformRequestQuery) (*ListofPlatforms, error)
type GetOnePlatformsHandler func (context.Context, *PlatformRequestParams) (*PlatformDetails, error)
type CreatePlatformsHandler func (context.Context, *PlatformRequestPayload) error
type UpdatePlatformsHandler func (context.Context, *PlatformRequestParams, *PlatformRequestPayload) error
type DeletePlatformsHandler func (context.Context, *PlatformRequestParams) error

func HandleGetAllPlatforms (handler GetAllPlatformsHandler) echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		ctx := c.Request().Context()
		query := &PlatformRequestQuery{}

		
		if err = c.Bind(query); err != nil {
			return err
		}
		
		if err = c.Validate(query); err!= nil {
			return err
		}

		data, err := handler(ctx,query)
		if err != nil {
			return err
		}

		return utils.WriteResponse(c, http.StatusOK,  data, "all platforms fetch successfully")
	}
}

func HandleGetOnePlatforms (handler GetOnePlatformsHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := &PlatformRequestParams{}

		
		if err := c.Bind(params); err != nil {
			return err
		}
		
		if err := c.Validate(params); err!= nil {
			return err
		}

		data, err := handler(ctx, params)
		if err != nil {
			return err
		}

		return utils.WriteResponse(c, 200, data, "platform fetch successfully")
	}
}

func HandleCreatePlatforms (handler CreatePlatformsHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		payload := &PlatformRequestPayload{}

		
		if err := c.Bind(payload); err != nil {
			return err
		}
		
		if err := c.Validate(payload); err!= nil {
			return err
		}

		err := handler(ctx, payload)
		if err != nil {
			return err
		}

		return utils.WriteResponse(c, http.StatusCreated, nil, "new platform successfully added")
	}
}

func HandleUpdatePlatforms (handler UpdatePlatformsHandler) echo.HandlerFunc{
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := &PlatformRequestParams{}
		payload := &PlatformRequestPayload{}

		if err := (&echo.DefaultBinder{}).BindPathParams(c, params); err != nil {
			return err
		}

		if err := c.Validate(params); err != nil {
			return err
		}

		if err := (&echo.DefaultBinder{}).BindBody(c, payload);err != nil {
			return err
		}

		if err := c.Validate(payload); err != nil {
			return err
		}

		if err := handler(ctx, params, payload); err != nil {
			return err
		}

		return utils.WriteResponse(c, 200, nil, "platform updated successfully")
	}
}

func HandleDeletePlatforms (handler DeletePlatformsHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := &PlatformRequestParams{}

		if err := c.Bind(params); err != nil {
			return err
		}

		if err := c.Validate(params); err != nil {
			return err
		}

		if err := handler(ctx, params); err != nil {
			return err
		}

		return utils.WriteResponse(c, 200, nil, "platform deleted successfully")
	}
}