package utils

func KeywordHelper(keyword *string) {
	if *keyword != "" {
		*keyword = "%" + *keyword + "%"
	} else {
		*keyword = "%"
	}
}