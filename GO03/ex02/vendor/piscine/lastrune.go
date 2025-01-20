package piscine

func LastRune(s string) rune {
	var res rune
	for _, r := range s {
		res = r
	}
	return res
}
