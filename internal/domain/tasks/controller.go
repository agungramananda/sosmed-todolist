package tasks

import "github.com/labstack/echo/v4"

type TasksController struct {
	svc TasksService
}

func NewController(svc TasksService) *TasksController {
	return &TasksController{
		svc: svc,
	}
}

const (
	tasksBasepath = "/tasks"
)

func (con *TasksController) Route(grp *echo.Group){
	subrouter := grp.Group(tasksBasepath)

	subrouter.GET("", HandleGetAllTasks(con.svc.GetAll))
	subrouter.GET("/:task_id", HandleGetOneTasks(con.svc.GetOne))
	subrouter.POST("",HandleCreateTasks(con.svc.Create))
	subrouter.PUT("/:task_id", HandleUpdateTasks(con.svc.Update))
	subrouter.DELETE("/:task_id", HandleDeleteTasks(con.svc.Delete))
}