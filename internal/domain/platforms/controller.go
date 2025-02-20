package platforms

import "github.com/labstack/echo/v4"

type PlatformsController struct {
	svc PlatformsService
}

func NewController(svc PlatformsService) *PlatformsController {
	return &PlatformsController{
		svc: svc,
	}
}

const (
	platformsBasepath = "/platforms"
)

func (con *PlatformsController) Route(grp *echo.Group){
	subrouter := grp.Group(platformsBasepath)

	subrouter.GET("", HandleGetAllPlatforms(con.svc.GetAll))
	subrouter.GET("/:platform_id", HandleGetOnePlatforms(con.svc.GetOne))
	subrouter.POST("",HandleCreatePlatforms(con.svc.Create))
	subrouter.PUT("/:platform_id", HandleUpdatePlatforms(con.svc.Update))
	subrouter.DELETE("/:platform_id", HandleDeletePlatforms(con.svc.Delete))
}