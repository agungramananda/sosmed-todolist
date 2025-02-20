package brands

import "github.com/labstack/echo/v4"

type BrandsController struct {
	svc BrandsService
}

func NewController(svc BrandsService) *BrandsController {
	return &BrandsController{
		svc: svc,
	}
}

const (
	brandsBasepath = "/brands"
)

func (con *BrandsController) Route(grp *echo.Group){
	subrouter := grp.Group(brandsBasepath)

	subrouter.GET("", HandleGetAllBrands(con.svc.GetAll))
	subrouter.GET("/:brand_id", HandleGetOneBrands(con.svc.GetOne))
	subrouter.POST("",HandleCreateBrands(con.svc.Create))
	subrouter.PUT("/:brand_id", HandleUpdateBrands(con.svc.Update))
	subrouter.DELETE("/:brand_id", HandleDeleteBrands(con.svc.Delete))
}