package piscine

func LastRune(s string) rune {
	lastrune := []rune(s)
	lens := len([]rune(s))
	return lastrune[lens-1]
}
