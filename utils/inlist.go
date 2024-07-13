package utils

func InList(key string, list []string) (bool, int) {
	for index, s := range list {
		if s == key {
			return true, index
		}
	}
	return false, -1
}
