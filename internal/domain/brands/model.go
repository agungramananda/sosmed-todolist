package brands

type Brands struct {
	BrandID int64  `db:"brand_id"`
	Brand   string `db:"brand"`
}