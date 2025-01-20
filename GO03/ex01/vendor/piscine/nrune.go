package piscine

import "ft"

furn NRune(s string, n int) rune {
	for i, r := range s {
		if i == n-1 {
			return r
		}
	}
	return 0
}