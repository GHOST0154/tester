package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	if IsPrime(nb) {
		return nb
	}
	nb++
	return FindNextPrime(nb)
}
