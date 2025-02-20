package platforms

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/agungramananda/sosmed-todolist/internal/common/exceptions"
	"github.com/agungramananda/sosmed-todolist/internal/utils"
	"github.com/jmoiron/sqlx"
)

var pgSquirell = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type PlatformsRepository interface {
	GetAll(context.Context, *PlatformRequestQuery) ([]*Platforms, error)
	Count(context.Context, *PlatformRequestQuery) (uint64, error)
	GetByID(context.Context, *PlatformRequestParams) (*Platforms, error)
	Add(context.Context, *PlatformRequestPayload) error
	Update(context.Context, *PlatformRequestPayload, *PlatformRequestParams) error
	Delete(context.Context, *PlatformRequestParams) error
}

type platformsRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) PlatformsRepository {
	return &platformsRepository{
		db: db,
	}
}

func (r *platformsRepository) GetAll(ctx context.Context, query *PlatformRequestQuery) (resp []*Platforms, err error) {
	keyword := query.Keyword
	utils.KeywordHelper(&keyword)

	stmt, args, _ := pgSquirell.Select("p.platform_id", "p.platform").From("platforms p").Where(squirrel.And{squirrel.Eq{"p.deleted_at": nil}, squirrel.ILike{"p.platform": keyword}}).Limit(query.Limit).Offset((query.Page - 1) * query.Limit).ToSql()

	resp = []*Platforms{}

	rows, err := r.db.QueryContext(ctx, stmt, args...)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		col := &Platforms{}

		if err = rows.Scan(&col.PlatformID, &col.Platform); err != nil {
			return resp, err
		}

		resp = append(resp, col)
	}

	return resp, nil
}

func (r *platformsRepository) GetByID(ctx context.Context, params *PlatformRequestParams) (resp *Platforms, err error) {
	stmt, args, _ := pgSquirell.Select("p.platform_id", "p.platform").From("platforms p").Where(squirrel.And{squirrel.Eq{"p.deleted_at": nil}, squirrel.Eq{"p.platform_id": params.PlatformID}}).ToSql()

	resp = &Platforms{}

	err = r.db.QueryRowxContext(ctx, stmt, args...).StructScan(resp)
	if err != nil && err != sql.ErrNoRows {
		return resp, err
	} else if err == sql.ErrNoRows {
		return resp, exceptions.NewNotFoundError("platforms not found")
	}

	return resp, nil
}

func (r *platformsRepository) Count(ctx context.Context, query *PlatformRequestQuery) (resp uint64, err error) {
	keyword := query.Keyword
	utils.KeywordHelper(&keyword)

	stmt, args, _ := pgSquirell.Select("count(platform_id)").From("platforms").Where(squirrel.And{squirrel.Eq{"deleted_at": nil}, squirrel.ILike{"platform":keyword}}).ToSql()

	err = r.db.QueryRowxContext(ctx, stmt, args...).Scan(&resp)
	if err != nil && err != sql.ErrNoRows {
		return resp, nil
	} else if err == sql.ErrNoRows {
		return 0, nil
	}

	return resp, nil
}

func (r *platformsRepository) Add(ctx context.Context, payload *PlatformRequestPayload) (err error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var stmt string
	var args []any

	stmt, args, _ = pgSquirell.Insert("platforms").Columns("platform").Values(payload.Platform).ToSql()

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *platformsRepository) Update(ctx context.Context, payload *PlatformRequestPayload, params *PlatformRequestParams) (err error) {
	tx, err := r.db.BeginTxx(ctx,nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var stmt string
	var args []any
	var count int64

	stmt, args, _ = pgSquirell.Select("count(*)").From("platforms").Where(squirrel.And{squirrel.Eq{"deleted_at":nil}, squirrel.Eq{"platform_id":params.PlatformID}}).ToSql()

	err = tx.QueryRowxContext(ctx, stmt, args...).Scan(&count)
	if err != nil && err != sql.ErrNoRows{
		return err
	} else if err == sql.ErrNoRows{
		return exceptions.NewInvariantError(err.Error())
	}

	stmt, args, _ = pgSquirell.Update("platforms").SetMap(map[string]interface{}{
		"platform":payload.Platform,
		"updated_at":squirrel.Expr("NOW()"),
	}).Where(squirrel.Eq{"platform_id":params.PlatformID}).ToSql()

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *platformsRepository) Delete(ctx context.Context, params *PlatformRequestParams) error {
	tx, err := r.db.BeginTxx(ctx, nil)

	if err != nil {
		return err
	}
	defer tx.Rollback()

	var stmt string
	var args []any
	var count int64

	stmt, args, _ = pgSquirell.Select("count(*)").From("platforms").Where(squirrel.And{squirrel.Eq{"deleted_at":nil}, squirrel.Eq{"platform_id":params.PlatformID}}).ToSql()

	err = tx.QueryRowxContext(ctx, stmt, args...).Scan(&count)
	if err != nil && err != sql.ErrNoRows{
		return err
	} else if err == sql.ErrNoRows{
		return exceptions.NewInvariantError(err.Error())
	}

	stmt, args, _ = pgSquirell.Update("platforms").SetMap(map[string]interface{}{
		"deleted_at":squirrel.Expr("NOW()"),
	}).Where(squirrel.Eq{"platform_id":params.PlatformID}).ToSql()

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}