package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	isWordFirst := true
	for i, r := range runes {
		if isWordFirst && r >= 'a' && r <= 'z' {
			runes[i] = r - 32
		} else if !isWordFirst && r >= 'A' && r <= 'Z' {
			runes[i] = r + 32
		}
		isWordFirst = r < 'A' || r > 'Z' && r < 'a' || r > 'z'
	}
	return string(runes)
}