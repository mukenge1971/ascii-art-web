package functions

func IsAski(s string) bool {
	for _, w := range s {
		if (w < 32 || w > 126) && (w != 10 && w != 13) {
			return false
		}
	}
	return true
}
