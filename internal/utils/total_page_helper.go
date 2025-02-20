package utils

func CountTotalPage(total_items uint64, limit uint64) uint64 {
	if limit == 0 {
		return 0
	}

	if total_items < limit {
		return 1
	}
	return total_items / limit
}