package piscine

import "ft"

func PrintReversealphabet() {
	for c := 'z'; c >= 'a'; c-- {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}