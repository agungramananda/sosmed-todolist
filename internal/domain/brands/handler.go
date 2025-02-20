package brands

import (
	"context"
	"net/http"

	"github.com/agungramananda/sosmed-todolist/internal/utils"
	"github.com/labstack/echo/v4"
)

type GetAllBrandsHandler func (context.Context, *BrandRequestQuery) (*ListofBrands, error)
type GetOneBrandsHandler func (context.Context, *BrandRequestParams) (*BrandDetails, error)
type CreateBrandsHandler func (context.Context, *BrandRequestPayload) error
type UpdateBrandsHandler func (context.Context, *BrandRequestParams, *BrandRequestPayload) error
type DeleteBrandsHandler func (context.Context, *BrandRequestParams) error

func HandleGetAllBrands (handler GetAllBrandsHandler) echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		ctx := c.Request().Context()
		query := &BrandRequestQuery{}

		
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

		return utils.WriteResponse(c, http.StatusOK,  data, "all brands fetch successfully")
	}
}

func HandleGetOneBrands (handler GetOneBrandsHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := &BrandRequestParams{}

		
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

		return utils.WriteResponse(c, 200, data, "brand fetch successfully")
	}
}

func HandleCreateBrands (handler CreateBrandsHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		payload := &BrandRequestPayload{}

		
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

		return utils.WriteResponse(c, http.StatusCreated, nil, "new brand successfully added")
	}
}

func HandleUpdateBrands (handler UpdateBrandsHandler) echo.HandlerFunc{
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := &BrandRequestParams{}
		payload := &BrandRequestPayload{}

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

		return utils.WriteResponse(c, 200, nil, "brand updated successfully")
	}
}

func HandleDeleteBrands (handler DeleteBrandsHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := &BrandRequestParams{}

		if err := c.Bind(params); err != nil {
			return err
		}

		if err := c.Validate(params); err != nil {
			return err
		}

		if err := handler(ctx, params); err != nil {
			return err
		}

		return utils.WriteResponse(c, 200, nil, "brand deleted successfully")
	}
}