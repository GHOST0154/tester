package piscine

func IsPrime(nb int) bool {
	if nb == 1 || nb <= 0 {
		return false
	}
	for v := 2; v*v <= nb; v++ {
		if nb%v == 0 {
			return false
		}
	}
	return true
}
