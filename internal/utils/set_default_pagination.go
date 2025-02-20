package utils

func SetDefaultPagination(limit *int, page *int) {
	if limit == nil || *limit == 0 {
		*limit = 25
	}
	if page == nil || *page == 0 {
		*page = 1
	}
}
