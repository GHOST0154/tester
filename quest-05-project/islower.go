package piscine

func IsLower(s string) bool {
	if s == "" {
		return false
	}
	for _, ch := range s {
		if ch < 'a' || ch > 'z' {
			return false
		}
	}
	return true
}
