package piscine

func UltimateDivMod(a *int, b *int) {
	result := *a / *b
	remainder := *a % *b

	*a = result
	*b = remainder
}
