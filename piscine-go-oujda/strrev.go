package piscine

func StrRev(s string) string {
	rev := ""
	for _, den := range s {
		rev = string(den) + rev
	}
	return rev
}
