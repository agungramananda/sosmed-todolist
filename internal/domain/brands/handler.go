package brands

import (
	"context"
	"net/http"

	"github.com/agungramananda/sosmed-todolist/internal/utils"
	"github.com/labstack/echo/v4"
)

type GetAllBrandsHandler func(context.Context, *BrandRequestQuery) (*ListofBrands, error)
type GetOneBrandsHandler func(context.Context, *BrandRequestParams) (*BrandDetails, error)
type CreateBrandsHandler func(context.Context, *BrandRequestPayload) error
type UpdateBrandsHandler func(context.Context, *BrandRequestParams, *BrandRequestPayload) error
type DeleteBrandsHandler func(context.Context, *BrandRequestParams) error

// Get All Brands godoc
//
//	@Summary	Get all brands
//	@Tags		Brand
//	@Produce	json
//	@Param		keyword	query		string	false	"Keyword to search"
//	@Param		limit	query		int		false	"Number of entities per page"
//	@Param		page	query		int		false	"Page number"
//	@Success	200		{object}	ListofBrands	"Successfully fetched all brands"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/brands [get]
func HandleGetAllBrands(handler GetAllBrandsHandler) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		ctx := c.Request().Context()
		query := &BrandRequestQuery{}

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

		return utils.WriteResponse(c, http.StatusOK, data, "All brands fetched successfully")
	}
}

// Get One Brand godoc
//
//	@Summary	Get a single brand by ID
//	@Tags		Brand
//	@Produce	json
//	@Param		id	path	string	true	"Brand ID"
//	@Success	200		{object}	BrandDetails	"Successfully fetched the brand"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	404		{object}	httpres.ErrorResponse	"Brand not found"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/brands/{id} [get]
func HandleGetOneBrands(handler GetOneBrandsHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := &BrandRequestParams{}

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

		return utils.WriteResponse(c, http.StatusOK, data, "Brand fetched successfully")
	}
}

// Create Brand godoc
//
//	@Summary	Create a new brand
//	@Tags		Brand
//	@Accept		json
//	@Produce	json
//	@Param		body	body	BrandRequestPayload	true	"Brand details"
//	@Success	201		{object}	httpres.BaseResponse	"Brand successfully created"
//	@Failure	400		{object}	httpres.ErrorResponse			"Bad request"
//	@Failure	500		{object}	httpres.ErrorResponse			"Internal server error"
//	@Router		/brands [post]
func HandleCreateBrands(handler CreateBrandsHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		payload := &BrandRequestPayload{}

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

		return utils.WriteResponse(c, http.StatusCreated, nil, "New brand successfully added")
	}
}

// Update Brand godoc
//
//	@Summary	Update an existing brand
//	@Tags		Brand
//	@Accept		json
//	@Produce	json
//	@Param		id		path	string					true	"Brand ID"
//	@Param		body	body	BrandRequestPayload	true	"Updated brand details"
//	@Success	200		{object}	httpres.BaseResponse	"Brand updated successfully"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	404		{object}	httpres.ErrorResponse	"Brand not found"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/brands/{id} [put]
func HandleUpdateBrands(handler UpdateBrandsHandler) echo.HandlerFunc {
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

		if err := (&echo.DefaultBinder{}).BindBody(c, payload); err != nil {
			return err
		}

		if err := c.Validate(payload); err != nil {
			return err
		}

		if err := handler(ctx, params, payload); err != nil {
			return err
		}

		return utils.WriteResponse(c, http.StatusOK, nil, "Brand updated successfully")
	}
}

// Delete Brand godoc
//
//	@Summary	Delete a brand by ID
//	@Tags		Brand
//	@Produce	json
//	@Param		id	path	string	true	"Brand ID"
//	@Success	200		{object}	httpres.BaseResponse	"Brand deleted successfully"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	404		{object}	httpres.ErrorResponse	"Brand not found"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/brands/{id} [delete]
func HandleDeleteBrands(handler DeleteBrandsHandler) echo.HandlerFunc {
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

		return utils.WriteResponse(c, http.StatusOK, nil, "Brand deleted successfully")
	}
}
