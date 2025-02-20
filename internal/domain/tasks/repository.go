package tasks

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/agungramananda/sosmed-todolist/internal/common/exceptions"
	"github.com/agungramananda/sosmed-todolist/internal/utils"
	"github.com/jmoiron/sqlx"
)

var pgSquirell = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type TasksRepository interface {
	GetAll(context.Context, *TaskRequestQuery) ([]*Tasks, error)
	Count(context.Context, *TaskRequestQuery) (uint64, error)
	GetByID(context.Context, *TaskRequestParams) (*Tasks, error)
	Add(context.Context, *TaskRequestPayload) error
	Update(context.Context, *TaskRequestPayload, *TaskRequestParams) error
	Delete(context.Context, *TaskRequestParams) error
}

type tasksRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) TasksRepository {
	return &tasksRepository{
		db: db,
	}
}

func (r *tasksRepository) GetAll(ctx context.Context, query *TaskRequestQuery) (resp []*Tasks, err error) {
	keyword := query.Keyword
	utils.KeywordHelper(&keyword)

	stmt, args, _ := pgSquirell.Select("t.task_id", "t.title", "t.brand_id", "b.brand", "t.platform_id", "p.platform", "t.due_date","t.payment","t.status").
						From("tasks t").
						LeftJoin("brands b on t.brand_id=b.brand_id").
						LeftJoin("platforms p on t.platform_id=p.platform_id").
						Where(squirrel.And{squirrel.Eq{"t.deleted_at": nil}, squirrel.ILike{"t.title": keyword}, squirrel.Eq{"b.deleted_at":nil}, squirrel.Eq{"p.deleted_at":nil}}).
						Limit(query.Limit).Offset((query.Page - 1) * query.Limit).ToSql()

	resp = []*Tasks{}

	rows, err := r.db.QueryContext(ctx, stmt, args...)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		col := &Tasks{}

		if err = rows.Scan(&col.TaskID, &col.Title, &col.BrandID, &col.Brand, &col.PlatformID, &col.Platform, &col.DueDate, &col.Payment, &col.Status); err != nil {
			return resp, err
		}

		resp = append(resp, col)
	}

	return resp, nil
}

func (r *tasksRepository) GetByID(ctx context.Context, params *TaskRequestParams) (resp *Tasks, err error) {
	stmt, args, _ := pgSquirell.Select("t.task_id", "t.title", "t.brand_id", "b.brand", "t.platform_id", "p.platform", "t.due_date","t.payment","t.status").
						From("tasks t").
						LeftJoin("brands b on t.brand_id=b.brand_id").
						LeftJoin("platforms p on t.platform_id=p.platform_id").
						Where(squirrel.And{squirrel.Eq{"t.deleted_at": nil}, squirrel.Eq{"t.task_id": params.TaskID}, squirrel.Eq{"b.deleted_at":nil}, squirrel.Eq{"p.deleted_at":nil}}).
						ToSql()

	resp = &Tasks{}

	err = r.db.QueryRowxContext(ctx, stmt, args...).StructScan(resp)
	if err != nil && err != sql.ErrNoRows {
		return resp, err
	} else if err == sql.ErrNoRows {
		return nil, exceptions.NewNotFoundError("tasks not found")
	}

	return resp, nil
}

func (r *tasksRepository) Count(ctx context.Context, query *TaskRequestQuery) (resp uint64, err error) {
	keyword := query.Keyword
	utils.KeywordHelper(&keyword)

	stmt, args, _ := pgSquirell.Select("count(t.task_id)").
						From("tasks t").
						LeftJoin("brands b on t.brand_id=b.brand_id").
						LeftJoin("platforms p on t.platform_id=p.platform_id").
						Where(squirrel.And{squirrel.Eq{"t.deleted_at": nil}, squirrel.Eq{"b.deleted_at":nil}, squirrel.Eq{"p.deleted_at":nil}}).
						ToSql()

	err = r.db.QueryRowxContext(ctx, stmt, args...).Scan(&resp)
	if err != nil && err != sql.ErrNoRows {
		return resp, nil
	} else if err == sql.ErrNoRows {
		return 0, nil
	}

	return resp, nil
}

func (r *tasksRepository) Add(ctx context.Context, payload *TaskRequestPayload) (err error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var stmt string
	var args []any
	var count int64

	stmt, args, _ = pgSquirell.Select("count(*)").From("brands").Where(squirrel.Eq{"brand_id": payload.BrandID, "deleted_at": nil}).ToSql()
	err = tx.QueryRowxContext(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return err
	} else if count == 0 {
		return exceptions.NewInvariantError("brand_id does not exist")
	}

	stmt, args, _ = pgSquirell.Select("count(*)").From("platforms").Where(squirrel.Eq{"platform_id": payload.PlatformID, "deleted_at": nil}).ToSql()
	err = tx.QueryRowxContext(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return err
	} else if count == 0 {
		return exceptions.NewInvariantError("platform_id does not exist")
	}

	stmt, args, _ = pgSquirell.Insert("tasks").Columns("title", "brand_id", "platform_id", "due_date", "payment", "status").Values(payload.Title, payload.BrandID, payload.PlatformID, payload.DueDate, payload.Payment, payload.Status).ToSql()

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *tasksRepository) Update(ctx context.Context, payload *TaskRequestPayload, params *TaskRequestParams) (err error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var stmt string
	var args []any
	var count int64

	stmt, args, _ = pgSquirell.Select("count(*)").From("brands").Where(squirrel.Eq{"brand_id": payload.BrandID, "deleted_at": nil}).ToSql()
	err = tx.QueryRowxContext(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return err
	} else if count == 0 {
		return exceptions.NewInvariantError("brand_id does not exist")
	}

	stmt, args, _ = pgSquirell.Select("count(*)").From("platforms").Where(squirrel.Eq{"platform_id": payload.PlatformID, "deleted_at": nil}).ToSql()
	err = tx.QueryRowxContext(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return err
	} else if count == 0 {
		return exceptions.NewInvariantError("platform_id does not exist")
	}

	stmt, args, _ = pgSquirell.Select("count(*)").From("tasks t").
		LeftJoin("brands b on t.brand_id=b.brand_id").
		LeftJoin("platforms p on t.platform_id=p.platform_id").
		Where(squirrel.And{squirrel.Eq{"t.deleted_at": nil}, squirrel.Eq{"b.deleted_at": nil}, squirrel.Eq{"p.deleted_at": nil}, squirrel.Eq{"t.task_id": params.TaskID}}).
		ToSql()

	err = tx.QueryRowxContext(ctx, stmt, args...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return err
	} else if err == sql.ErrNoRows {
		return exceptions.NewInvariantError(err.Error())
	}

	stmt, args, _ = pgSquirell.Update("tasks").SetMap(map[string]interface{}{
		"title":       payload.Title,
		"brand_id":    payload.BrandID,
		"platform_id": payload.PlatformID,
		"due_date":    payload.DueDate,
		"payment":     payload.Payment,
		"status":      payload.Status,
		"updated_at":  squirrel.Expr("NOW()"),
	}).Where(squirrel.Eq{"task_id": params.TaskID}).ToSql()

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *tasksRepository) Delete(ctx context.Context, params *TaskRequestParams) error {
	tx, err := r.db.BeginTxx(ctx, nil)

	if err != nil {
		return err
	}
	defer tx.Rollback()

	var stmt string
	var args []any
	var count int64

	stmt, args, _ = pgSquirell.Select("count(*)").From("tasks t").
					LeftJoin("brands b on t.brand_id=b.brand_id").
					LeftJoin("platforms p on t.platform_id=p.platform_id").
					Where(squirrel.And{squirrel.Eq{"t.deleted_at": nil}, squirrel.Eq{"b.deleted_at":nil}, squirrel.Eq{"p.deleted_at":nil}, squirrel.Eq{"t.task_id":params.TaskID}}).
					ToSql()

	err = tx.QueryRowxContext(ctx, stmt, args...).Scan(&count)
	if err != nil && err != sql.ErrNoRows{
		return err
	} else if err == sql.ErrNoRows{
		return exceptions.NewInvariantError(err.Error())
	}

	stmt, args, _ = pgSquirell.Update("tasks").SetMap(map[string]interface{}{
		"deleted_at":squirrel.Expr("NOW()"),
	}).Where(squirrel.Eq{"task_id":params.TaskID}).ToSql()

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}