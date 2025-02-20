package brands

import "github.com/agungramananda/sosmed-todolist/internal/common/httpres"

type BrandRequestParams struct {
	BrandID string `param:"brand_id" validate:"required"`
}

type BrandRequestPayload struct {
	BrandID int64  `json:"-"`
	Brand   string `json:"brand" validate:"required,max=200,min=1"`
}

type BrandRequestQuery struct {
	Keyword string `query:"keyword" validate:"omitempty,max=100"`
	Limit   uint64 `query:"limit" validate:"omitempty,min=1,max=100"`
	Page    uint64 `query:"page" validate:"omitempty,min=1"`
}

type BrandDetails struct {
	BrandID int64  `json:"brand_id"`
	Brand   string `json:"brand"`
}

type ListofBrands struct {
	Brands []*BrandDetails `json:"brands"`
	Meta   httpres.ListPagination `json:"meta"`
}