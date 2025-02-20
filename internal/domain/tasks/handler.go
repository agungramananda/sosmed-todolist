package tasks

import (
	"context"
	"net/http"

	"github.com/agungramananda/sosmed-todolist/internal/utils"
	"github.com/labstack/echo/v4"
)

type GetAllTasksHandler func (context.Context, *TaskRequestQuery) (*ListofTasks, error)
type GetOneTasksHandler func (context.Context, *TaskRequestParams) (*TaskDetails, error)
type CreateTasksHandler func (context.Context, *TaskRequestPayload) error
type UpdateTasksHandler func (context.Context, *TaskRequestParams, *TaskRequestPayload) error
type DeleteTasksHandler func (context.Context, *TaskRequestParams) error

func HandleGetAllTasks (handler GetAllTasksHandler) echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		ctx := c.Request().Context()
		query := &TaskRequestQuery{}

		
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

		return utils.WriteResponse(c, http.StatusOK,  data, "all tasks fetch successfully")
	}
}

func HandleGetOneTasks (handler GetOneTasksHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := &TaskRequestParams{}

		
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

		return utils.WriteResponse(c, 200, data, "task fetch successfully")
	}
}

func HandleCreateTasks (handler CreateTasksHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		payload := &TaskRequestPayload{}

		
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

		return utils.WriteResponse(c, http.StatusCreated, nil, "new task successfully added")
	}
}

func HandleUpdateTasks (handler UpdateTasksHandler) echo.HandlerFunc{
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := &TaskRequestParams{}
		payload := &TaskRequestPayload{}

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

		return utils.WriteResponse(c, 200, nil, "task updated successfully")
	}
}

func HandleDeleteTasks (handler DeleteTasksHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := &TaskRequestParams{}

		if err := c.Bind(params); err != nil {
			return err
		}

		if err := c.Validate(params); err != nil {
			return err
		}

		if err := handler(ctx, params); err != nil {
			return err
		}

		return utils.WriteResponse(c, 200, nil, "task deleted successfully")
	}
}