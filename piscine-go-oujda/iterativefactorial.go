package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 {
		return 0
	}
	if nb == 1 || nb == 0 {
		return 1
	}
	for i := 2; i <= nb; i++ {
		result *= i
		if result < 0 {
			return 0
		}

	}
	return result
}
