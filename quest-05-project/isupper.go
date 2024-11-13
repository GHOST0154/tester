package piscine

func IsUpper(s string) bool {
	if s == "" {
		return false
	}
	for _, ch := range s {
		if ch < 'A' || ch > 'Z' {
			return false
		}
	}
	return true
}
