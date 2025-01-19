package piscine

import "ft"

func RecursivePower(nb int, power int) int {
	if n <= 0 {
		return 1
	}
	return nb * RecursivePower(nb, power -1)
}

