package piscine

import "ft"

func IsNegative(nb int) {
	if (nb < 0)
		PrintRune('T')
	else
		PrintRune('F')
	PrintRune('\n')
}