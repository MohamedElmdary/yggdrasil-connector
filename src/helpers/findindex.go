package helpers

func FindIndex(list []string, item string) int {
	idx := -1
	for i, p := range list {
		if p == item {
			idx = i
			break
		}
	}
	return idx
}
