package tasks

import "github.com/agungramananda/sosmed-todolist/internal/common/httpres"

type TaskRequestParams struct {
	TaskID string `param:"task_id" validate:"required"`
}

type TaskRequestPayload struct {
	TaskID     int64  `json:"-"`
	Title      string `json:"title" validate:"required"`
	BrandID    int64  `json:"brand_id" validate:"omitempty,min=1"`
	PlatformID int64  `json:"platform_id" validate:"omitempty,min=1"`
	DueDate    string `json:"due_date" validate:"required,datetime=2006-01-02"`
	Payment    int64  `json:"payment" validate:"required"`
	Status     string `json:"status" validate:"required,oneof='Pending' 'Completed' 'Scheduled'"`
}

type TaskRequestQuery struct {
	Keyword string `query:"keyword" validate:"omitempty,max=100"`
	Limit   uint64 `query:"limit" validate:"omitempty,min=1,max=100"`
	Page    uint64 `query:"page" validate:"omitempty,min=1"`
}

type TaskDetails struct {
	TaskID int64  `json:"task_id"`
	Title      string `json:"title"`
	BrandID    int64  `json:"brand_id"`
	Brand      string `json:"brand"`
	PlatformID int64  `json:"platform_id"`
	Platform   string `json:"platform"`
	DueDate    string `json:"due_date"`
	Payment    string  `json:"payment"`
	Status     string `json:"status"`
}

type ListofTasks struct {
	Tasks []*TaskDetails `json:"tasks"`
	Meta   httpres.ListPagination `json:"meta"`
}