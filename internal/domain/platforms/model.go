package platforms

type Platforms struct {
	PlatformID int64  `db:"platform_id"`
	Platform   string `db:"platform"`
}