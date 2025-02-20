package tasks

import "time"

type Tasks struct {
	TaskID 		int64     	`db:"task_id"`
	Title 		string    	`db:"title"`
	BrandID 	int64     	`db:"brand_id"`
	Brand		string 		`db:"brand"`
	PlatformID 	int64     	`db:"platform_id"`
	Platform	string 		`db:"platform"`
	DueDate 	time.Time 	`db:"due_date"`
	Payment 	string 	 	`db:"payment"`
	Status 		string		`db:"status"`
}