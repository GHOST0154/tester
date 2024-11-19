package piscine

func AppendRange(min, max int) []int {
	base := []int{}
	if min >= max {
		return []int(nil)
	}
	for min < max {
		base = append(base, min)
		min++
	}
	return base
}
