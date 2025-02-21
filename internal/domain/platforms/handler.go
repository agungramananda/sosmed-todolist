package platforms

import (
	"context"
	"net/http"

	"github.com/agungramananda/sosmed-todolist/internal/utils"
	"github.com/labstack/echo/v4"
)

type GetAllPlatformsHandler func(context.Context, *PlatformRequestQuery) (*ListofPlatforms, error)
type GetOnePlatformsHandler func(context.Context, *PlatformRequestParams) (*PlatformDetails, error)
type CreatePlatformsHandler func(context.Context, *PlatformRequestPayload) error
type UpdatePlatformsHandler func(context.Context, *PlatformRequestParams, *PlatformRequestPayload) error
type DeletePlatformsHandler func(context.Context, *PlatformRequestParams) error

// Get All Platforms godoc
//
//	@Summary	Get all platforms
//	@Tags		Platform
//	@Produce	json
//	@Param		keyword	query		string	false	"Keyword to search"
//	@Param		limit	query		int		false	"Number of entities per page"
//	@Param		page	query		int		false	"Page number"
//	@Success	200		{object}	ListofPlatforms	"Successfully fetched all platforms"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/platforms [get]
func HandleGetAllPlatforms(handler GetAllPlatformsHandler) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		ctx := c.Request().Context()
		query := &PlatformRequestQuery{}

		if err = c.Bind(query); err != nil {
			return err
		}

		if err = c.Validate(query); err != nil {
			return err
		}

		data, err := handler(ctx, query)
		if err != nil {
			return err
		}

		return utils.WriteResponse(c, http.StatusOK, data, "all platforms fetch successfully")
	}
}

// Get One Platform godoc
//
//	@Summary	Get a single platform by ID
//	@Tags		Platform
//	@Produce	json
//	@Param		id	path	string	true	"Platform ID"
//	@Success	200		{object}	PlatformDetails	"Successfully fetched the platform"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	404		{object}	httpres.ErrorResponse	"Platform not found"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/platforms/{id} [get]
func HandleGetOnePlatforms(handler GetOnePlatformsHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := &PlatformRequestParams{}

		if err := c.Bind(params); err != nil {
			return err
		}

		if err := c.Validate(params); err != nil {
			return err
		}

		data, err := handler(ctx, params)
		if err != nil {
			return err
		}

		return utils.WriteResponse(c, 200, data, "platform fetch successfully")
	}
}

// Create Platform godoc
//
//	@Summary	Create a new platform
//	@Tags		Platform
//	@Accept		json
//	@Produce	json
//	@Param		body	body	PlatformRequestPayload	true	"Platform details"
//	@Success	201		{object}	httpres.BaseResponse	"Platform successfully created"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/platforms [post]
func HandleCreatePlatforms(handler CreatePlatformsHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		payload := &PlatformRequestPayload{}

		if err := c.Bind(payload); err != nil {
			return err
		}

		if err := c.Validate(payload); err != nil {
			return err
		}

		err := handler(ctx, payload)
		if err != nil {
			return err
		}

		return utils.WriteResponse(c, http.StatusCreated, nil, "new platform successfully added")
	}
}

// Update Platform godoc
//
//	@Summary	Update an existing platform
//	@Tags		Platform
//	@Accept		json
//	@Produce	json
//	@Param		id		path	string					true	"Platform ID"
//	@Param		body	body	PlatformRequestPayload	true	"Updated platform details"
//	@Success	200		{object}	httpres.BaseResponse	"Platform updated successfully"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	404		{object}	httpres.ErrorResponse	"Platform not found"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/platforms/{id} [put]
func HandleUpdatePlatforms(handler UpdatePlatformsHandler) echo.HandlerFunc {
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

		if err := (&echo.DefaultBinder{}).BindBody(c, payload); err != nil {
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

// Delete Platform godoc
//
//	@Summary	Delete a platform by ID
//	@Tags		Platform
//	@Produce	json
//	@Param		id	path	string	true	"Platform ID"
//	@Success	200		{object}	httpres.BaseResponse	"Platform deleted successfully"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	404		{object}	httpres.ErrorResponse	"Platform not found"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/platforms/{id} [delete]
func HandleDeletePlatforms(handler DeletePlatformsHandler) echo.HandlerFunc {
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