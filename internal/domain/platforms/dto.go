package platforms

import "github.com/agungramananda/sosmed-todolist/internal/common/httpres"

type PlatformRequestParams struct {
	PlatformID string `param:"platform_id" validate:"required"`
}

type PlatformRequestPayload struct {
	PlatformID int64  `json:"-"`
	Platform   string `json:"platform" validate:"required,max=200,min=1"`
}

type PlatformRequestQuery struct {
	Keyword string `query:"keyword" validate:"omitempty,max=100"`
	Limit   uint64 `query:"limit" validate:"omitempty,min=1,max=100"`
	Page    uint64 `query:"page" validate:"omitempty,min=1"`
}

type PlatformDetails struct {
	PlatformID int64  `json:"platform_id"`
	Platform   string `json:"platform"`
}

type ListofPlatforms struct {
	Platforms []*PlatformDetails `json:"platforms"`
	Meta   httpres.ListPagination `json:"meta"`
}