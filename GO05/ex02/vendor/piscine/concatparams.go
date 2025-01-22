package piscine

func ConcatParams(args []string) string {
	var result string
	for i, s := range args {
		if i != 0 {
			result += "\n"
		}
		result += s
	}
	return result
}