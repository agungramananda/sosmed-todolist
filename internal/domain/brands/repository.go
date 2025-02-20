package brands

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/agungramananda/sosmed-todolist/internal/common/exceptions"
	"github.com/agungramananda/sosmed-todolist/internal/utils"
	"github.com/jmoiron/sqlx"
)

var pgSquirell = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type BrandsRepository interface {
	GetAll(context.Context, *BrandRequestQuery) ([]*Brands, error)
	Count(context.Context, *BrandRequestQuery) (uint64, error)
	GetByID(context.Context, *BrandRequestParams) (*Brands, error)
	Add(context.Context, *BrandRequestPayload) error
	Update(context.Context, *BrandRequestPayload, *BrandRequestParams) error
	Delete(context.Context, *BrandRequestParams) error
}

type brandsRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) BrandsRepository {
	return &brandsRepository{
		db: db,
	}
}

func (r *brandsRepository) GetAll(ctx context.Context, query *BrandRequestQuery) (resp []*Brands, err error) {
	keyword := query.Keyword
	utils.KeywordHelper(&keyword)

	stmt, args, _ := pgSquirell.Select("b.brand_id", "b.brand").From("brands b").Where(squirrel.And{squirrel.Eq{"b.deleted_at": nil}, squirrel.ILike{"b.brand": keyword}}).Limit(query.Limit).Offset((query.Page - 1) * query.Limit).ToSql()

	resp = []*Brands{}

	rows, err := r.db.QueryContext(ctx, stmt, args...)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		col := &Brands{}

		if err = rows.Scan(&col.BrandID, &col.Brand); err != nil {
			return resp, err
		}

		resp = append(resp, col)
	}

	return resp, nil
}

func (r *brandsRepository) GetByID(ctx context.Context, params *BrandRequestParams) (resp *Brands, err error) {
	stmt, args, _ := pgSquirell.Select("b.brand_id", "b.brand").From("brands b").Where(squirrel.And{squirrel.Eq{"b.deleted_at": nil}, squirrel.Eq{"b.brand_id": params.BrandID}}).ToSql()

	resp = &Brands{}

	err = r.db.QueryRowxContext(ctx, stmt, args...).StructScan(resp)
	if err != nil && err != sql.ErrNoRows {
		return resp, err
	} else if err == sql.ErrNoRows {
		return resp, exceptions.NewNotFoundError("brands not found")
	}

	return resp, nil
}

func (r *brandsRepository) Count(ctx context.Context, query *BrandRequestQuery) (resp uint64, err error) {
	keyword := query.Keyword
	utils.KeywordHelper(&keyword)

	stmt, args, _ := pgSquirell.Select("count(brand_id)").From("brands").Where(squirrel.And{squirrel.Eq{"deleted_at": nil}, squirrel.Eq{"brand":keyword}}).ToSql()

	err = r.db.QueryRowxContext(ctx, stmt, args...).Scan(&resp)
	if err != nil && err != sql.ErrNoRows {
		return resp, nil
	} else if err == sql.ErrNoRows {
		return 0, nil
	}

	return resp, nil
}

func (r *brandsRepository) Add(ctx context.Context, payload *BrandRequestPayload) (err error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var stmt string
	var args []any

	stmt, args, _ = pgSquirell.Insert("brands").Columns("brand").Values(payload.Brand).ToSql()

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *brandsRepository) Update(ctx context.Context, payload *BrandRequestPayload, params *BrandRequestParams) (err error) {
	tx, err := r.db.BeginTxx(ctx,nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var stmt string
	var args []any
	var count int64

	stmt, args, _ = pgSquirell.Select("count(*)").From("brands").Where(squirrel.And{squirrel.Eq{"deleted_at":nil}, squirrel.Eq{"brand_id":params.BrandID}}).ToSql()

	err = tx.QueryRowxContext(ctx, stmt, args...).Scan(&count)
	if err != nil && err != sql.ErrNoRows{
		return err
	} else if err == sql.ErrNoRows{
		return exceptions.NewInvariantError(err.Error())
	}

	stmt, args, _ = pgSquirell.Update("brands").SetMap(map[string]interface{}{
		"brand":payload.Brand,
		"updated_at":squirrel.Expr("NOW()"),
	}).Where(squirrel.Eq{"brand_id":params.BrandID}).ToSql()

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *brandsRepository) Delete(ctx context.Context, params *BrandRequestParams) error {
	tx, err := r.db.BeginTxx(ctx, nil)

	if err != nil {
		return err
	}
	defer tx.Rollback()

	var stmt string
	var args []any
	var count int64

	stmt, args, _ = pgSquirell.Select("count(*)").From("brands").Where(squirrel.And{squirrel.Eq{"deleted_at":nil}, squirrel.Eq{"brand_id":params.BrandID}}).ToSql()

	err = tx.QueryRowxContext(ctx, stmt, args...).Scan(&count)
	if err != nil && err != sql.ErrNoRows{
		return err
	} else if err == sql.ErrNoRows{
		return exceptions.NewInvariantError(err.Error())
	}

	stmt, args, _ = pgSquirell.Update("brands").SetMap(map[string]interface{}{
		"deleted_at":squirrel.Expr("NOW()"),
	}).Where(squirrel.Eq{"brand_id":params.BrandID}).ToSql()

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}