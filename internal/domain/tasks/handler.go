package tasks

import (
	"context"
	"net/http"

	"github.com/agungramananda/sosmed-todolist/internal/utils"
	"github.com/labstack/echo/v4"
)

type GetAllTasksHandler func(context.Context, *TaskRequestQuery) (*ListofTasks, error)
type GetOneTasksHandler func(context.Context, *TaskRequestParams) (*TaskDetails, error)
type CreateTasksHandler func(context.Context, *TaskRequestPayload) error
type UpdateTasksHandler func(context.Context, *TaskRequestParams, *TaskRequestPayload) error
type DeleteTasksHandler func(context.Context, *TaskRequestParams) error

// Get All Tasks godoc
//
//	@Summary	Get all tasks
//	@Tags		Task
//	@Produce	json
//	@Param		keyword	query		string	false	"Keyword to search"
//	@Param		limit	query		int		false	"Number of entities per page"
//	@Param		page	query		int		false	"Page number"
//	@Success	200		{object}	ListofTasks	"Successfully fetched all tasks"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/tasks [get]
func HandleGetAllTasks(handler GetAllTasksHandler) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		ctx := c.Request().Context()
		query := &TaskRequestQuery{}

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

		return utils.WriteResponse(c, http.StatusOK, data, "All tasks fetched successfully")
	}
}

// Get One Task godoc
//
//	@Summary	Get a single task by ID
//	@Tags		Task
//	@Produce	json
//	@Param		id	path	string	true	"Task ID"
//	@Success	200		{object}	TaskDetails	"Successfully fetched the task"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	404		{object}	httpres.ErrorResponse	"Task not found"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/tasks/{id} [get]
func HandleGetOneTasks(handler GetOneTasksHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := &TaskRequestParams{}

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

		return utils.WriteResponse(c, http.StatusOK, data, "Task fetched successfully")
	}
}

// Create Task godoc
//
//	@Summary	Create a new task
//	@Tags		Task
//	@Accept		json
//	@Produce	json
//	@Param		body	body	TaskRequestPayload	true	"Task details"
//	@Success	201		{object}	httpres.BaseResponse	"Task successfully created"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/tasks [post]
func HandleCreateTasks(handler CreateTasksHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		payload := &TaskRequestPayload{}

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

		return utils.WriteResponse(c, http.StatusCreated, nil, "New task successfully added")
	}
}

// Update Task godoc
//
//	@Summary	Update an existing task
//	@Tags		Task
//	@Accept		json
//	@Produce	json
//	@Param		id		path	string				true	"Task ID"
//	@Param		body	body	TaskRequestPayload	true	"Updated task details"
//	@Success	200		{object}	httpres.BaseResponse	"Task updated successfully"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	404		{object}	httpres.ErrorResponse	"Task not found"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/tasks/{id} [put]
func HandleUpdateTasks(handler UpdateTasksHandler) echo.HandlerFunc {
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

		if err := (&echo.DefaultBinder{}).BindBody(c, payload); err != nil {
			return err
		}

		if err := c.Validate(payload); err != nil {
			return err
		}

		if err := handler(ctx, params, payload); err != nil {
			return err
		}

		return utils.WriteResponse(c, http.StatusOK, nil, "Task updated successfully")
	}
}

// Delete Task godoc
//
//	@Summary	Delete a task by ID
//	@Tags		Task
//	@Produce	json
//	@Param		id	path	string	true	"Task ID"
//	@Success	200		{object}	httpres.BaseResponse	"Task deleted successfully"
//	@Failure	400		{object}	httpres.ErrorResponse	"Bad request"
//	@Failure	404		{object}	httpres.ErrorResponse	"Task not found"
//	@Failure	500		{object}	httpres.ErrorResponse	"Internal server error"
//	@Router		/tasks/{id} [delete]
func HandleDeleteTasks(handler DeleteTasksHandler) echo.HandlerFunc {
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

		return utils.WriteResponse(c, http.StatusOK, nil, "Task deleted successfully")
	}
}