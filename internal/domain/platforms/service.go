package platforms

import (
	"context"

	"github.com/agungramananda/sosmed-todolist/internal/common/httpres"
	"github.com/agungramananda/sosmed-todolist/internal/utils"
)

type PlatformsService interface {
	GetAll(context.Context, *PlatformRequestQuery) (*ListofPlatforms, error)
	GetOne(context.Context, *PlatformRequestParams) (*PlatformDetails, error)
	Create(context.Context, *PlatformRequestPayload) error
	Update(context.Context, *PlatformRequestParams, *PlatformRequestPayload) error
	Delete(context.Context, *PlatformRequestParams) error
}

type platformsService struct {
	repo      PlatformsRepository
}

func NewService(r PlatformsRepository) *platformsService {
	return &platformsService{repo: r}
}

func (svc platformsService) GetAll(ctx context.Context, query *PlatformRequestQuery) (listOfPlatforms *ListofPlatforms, err error) {
	limit := int(query.Limit)
	page := int(query.Page)
	utils.SetDefaultPagination(&limit, &page)

	repoQuery := &PlatformRequestQuery{
		Keyword: query.Keyword,
		Limit: uint64(limit),
		Page: uint64(page),
	}


	listOfPlatforms = &ListofPlatforms{
		Platforms: []*PlatformDetails{},
		Meta: httpres.ListPagination{
			Limit:		repoQuery.Limit,
			Page: 		repoQuery.Page,
			TotalPage: 	0,
		},
	}

	platforms, err := svc.repo.GetAll(ctx, repoQuery)
	if err != nil {
		return &ListofPlatforms{}, err
	}

	for _, platform := range platforms {
		listOfPlatforms.Platforms = append(listOfPlatforms.Platforms, &PlatformDetails{
			PlatformID: platform.PlatformID,
			Platform:   platform.Platform,
		})
	}

	total_items, err := svc.repo.Count(ctx, repoQuery)
	if err != nil {
		return &ListofPlatforms{}, err
	}

	listOfPlatforms.Meta.TotalPage = utils.CountTotalPage(total_items,repoQuery.Limit)

	return listOfPlatforms, nil
}

func (svc *platformsService) GetOne(ctx context.Context, params *PlatformRequestParams) (platformDetails *PlatformDetails, err error){
	platform, err := svc.repo.GetByID(ctx, params)
	if err != nil {
		return platformDetails, err
	}

	platformDetails = &PlatformDetails{
		PlatformID: platform.PlatformID,
		Platform:   platform.Platform,
	}

	return platformDetails, nil
}

func (svc *platformsService) Create(ctx context.Context, payload *PlatformRequestPayload) (err error) {
	err = svc.repo.Add(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (svc *platformsService) Update(ctx context.Context, params *PlatformRequestParams, payload *PlatformRequestPayload) (err error){
	err = svc.repo.Update(ctx,payload,params)
	if err != nil {
		return err
	}

	return nil
}

func (svc *platformsService) Delete(ctx context.Context, params *PlatformRequestParams) (err error){
	err = svc.repo.Delete(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
