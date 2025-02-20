package brands

import (
	"context"

	"github.com/agungramananda/sosmed-todolist/internal/common/httpres"
	"github.com/agungramananda/sosmed-todolist/internal/utils"
)

type BrandsService interface {
	GetAll(context.Context, *BrandRequestQuery) (*ListofBrands, error)
	GetOne(context.Context, *BrandRequestParams) (*BrandDetails, error)
	Create(context.Context, *BrandRequestPayload) error
	Update(context.Context, *BrandRequestParams, *BrandRequestPayload) error
	Delete(context.Context, *BrandRequestParams) error
}

type brandsService struct {
	repo      BrandsRepository
}

func NewService(r BrandsRepository) *brandsService {
	return &brandsService{repo: r}
}

func (svc brandsService) GetAll(ctx context.Context, query *BrandRequestQuery) (listOfBrands *ListofBrands, err error) {
	limit := int(query.Limit)
	page := int(query.Page)
	utils.SetDefaultPagination(&limit, &page)

	repoQuery := &BrandRequestQuery{
		Keyword: query.Keyword,
		Limit: uint64(limit),
		Page: uint64(page),
	}


	listOfBrands = &ListofBrands{
		Brands: []*BrandDetails{},
		Meta: httpres.ListPagination{
			Limit:		repoQuery.Limit,
			Page: 		repoQuery.Page,
			TotalPage: 	0,
		},
	}

	brands, err := svc.repo.GetAll(ctx, repoQuery)
	if err != nil {
		return &ListofBrands{}, err
	}

	for _, brand := range brands {
		listOfBrands.Brands = append(listOfBrands.Brands, &BrandDetails{
			BrandID: brand.BrandID,
			Brand:   brand.Brand,
		})
	}

	total_items, err := svc.repo.Count(ctx, repoQuery)
	if err != nil {
		return &ListofBrands{}, err
	}

	listOfBrands.Meta.TotalPage = utils.CountTotalPage(total_items,repoQuery.Limit)

	return listOfBrands, nil
}

func (svc *brandsService) GetOne(ctx context.Context, params *BrandRequestParams) (brandDetails *BrandDetails, err error){
	brand, err := svc.repo.GetByID(ctx, params)
	if err != nil {
		return brandDetails, err
	}

	brandDetails = &BrandDetails{
		BrandID: brand.BrandID,
		Brand:   brand.Brand,
	}

	return brandDetails, nil
}

func (svc *brandsService) Create(ctx context.Context, payload *BrandRequestPayload) (err error) {
	err = svc.repo.Add(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (svc *brandsService) Update(ctx context.Context, params *BrandRequestParams, payload *BrandRequestPayload) (err error){
	err = svc.repo.Update(ctx,payload,params)
	if err != nil {
		return err
	}

	return nil
}

func (svc *brandsService) Delete(ctx context.Context, params *BrandRequestParams) (err error){
	err = svc.repo.Delete(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
