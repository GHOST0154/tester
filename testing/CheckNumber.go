package piscine

func CheckNumber(arg string) bool {
	for _, i :=  range arg {
		if i>= '0' && i<= '9' {
			return true
		}
	}
	return false
}