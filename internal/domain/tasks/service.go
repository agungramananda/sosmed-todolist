package tasks

import (
	"context"

	"github.com/agungramananda/sosmed-todolist/internal/common/httpres"
	"github.com/agungramananda/sosmed-todolist/internal/utils"
)

type TasksService interface {
	GetAll(context.Context, *TaskRequestQuery) (*ListofTasks, error)
	GetOne(context.Context, *TaskRequestParams) (*TaskDetails, error)
	Create(context.Context, *TaskRequestPayload) error
	Update(context.Context, *TaskRequestParams, *TaskRequestPayload) error
	Delete(context.Context, *TaskRequestParams) error
}

type tasksService struct {
	repo      TasksRepository
}

func NewService(r TasksRepository) *tasksService {
	return &tasksService{repo: r}
}

func (svc tasksService) GetAll(ctx context.Context, query *TaskRequestQuery) (listOfTasks *ListofTasks, err error) {
	limit := int(query.Limit)
	page := int(query.Page)
	utils.SetDefaultPagination(&limit, &page)

	repoQuery := &TaskRequestQuery{
		Keyword: query.Keyword,
		Limit: uint64(limit),
		Page: uint64(page),
	}


	listOfTasks = &ListofTasks{
		Tasks: []*TaskDetails{},
		Meta: httpres.ListPagination{
			Limit:		repoQuery.Limit,
			Page: 		repoQuery.Page,
			TotalPage: 	0,
		},
	}

	tasks, err := svc.repo.GetAll(ctx, repoQuery)
	if err != nil {
		return &ListofTasks{}, err
	}

	for _, task := range tasks {
		listOfTasks.Tasks = append(listOfTasks.Tasks, &TaskDetails{
			TaskID:     task.TaskID,
			Title:      task.Title,
			BrandID:    task.BrandID,
			Brand:      task.Brand,
			PlatformID: task.PlatformID,
			Platform:   task.Platform,
			DueDate:    task.DueDate.Format("2006-01-02"),
			Payment:    task.Payment,
			Status:     task.Status,
		})
	}

	total_items, err := svc.repo.Count(ctx, repoQuery)
	if err != nil {
		return &ListofTasks{}, err
	}

	listOfTasks.Meta.TotalPage = utils.CountTotalPage(total_items,repoQuery.Limit)

	return listOfTasks, nil
}

func (svc *tasksService) GetOne(ctx context.Context, params *TaskRequestParams) (taskDetails *TaskDetails, err error){
	task, err := svc.repo.GetByID(ctx, params)
	if err != nil {
		return taskDetails, err
	}

	taskDetails = &TaskDetails{
		TaskID:     task.TaskID,
		Title:      task.Title,
		BrandID:    task.BrandID,
		Brand:      task.Brand,
		PlatformID: task.PlatformID,
		Platform:   task.Platform,
		DueDate:    task.DueDate.Format("2006-01-02"),
		Payment:    task.Payment,
		Status:     task.Status,
	}

	return taskDetails, nil
}

func (svc *tasksService) Create(ctx context.Context, payload *TaskRequestPayload) (err error) {
	err = svc.repo.Add(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (svc *tasksService) Update(ctx context.Context, params *TaskRequestParams, payload *TaskRequestPayload) (err error){
	err = svc.repo.Update(ctx,payload,params)
	if err != nil {
		return err
	}

	return nil
}

func (svc *tasksService) Delete(ctx context.Context, params *TaskRequestParams) (err error){
	err = svc.repo.Delete(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
