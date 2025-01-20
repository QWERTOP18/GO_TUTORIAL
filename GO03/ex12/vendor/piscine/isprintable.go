package piscine

func IsPrintable(str string) bool {
	for _, r := range str {
		if r < 32 {
			return false
		}
	}
	return true
}