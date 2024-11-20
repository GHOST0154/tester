package piscine

func IsPrintable(s string) bool {
	for _, i := range s {
		if (i <= 31){
			return false
		}
	}
	return true
}