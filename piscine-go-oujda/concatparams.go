package piscine

func ConcatParams(args []string) string {
	b := ""

	stop := len(args) - 1
	for i := 0; i < len(args); i++ {
		b += args[i]
		if i < stop {
			b += "\n"
		}
	}
	return b
}
