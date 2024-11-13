package piscine

func IsPrintable(s string) bool {
	for _, p := range s {
		if p < 32 || p > 126 {
			return true
		}
	}
	return false
}
