package httpres

type ListPagination struct {
	Limit     uint64 `json:"limit" example:"100"`
	Page      uint64 `json:"page" example:"1"`
	TotalPage uint64 `json:"total_page" example:"10"`
}

type BaseResponse struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}
