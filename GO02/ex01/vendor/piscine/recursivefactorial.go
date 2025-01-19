package piscine

import "ft"

func RecursiveFactorial(nb int) int {
	if nb <= 0 {
		return 1
	}
	return nb * RecursiveFactorial(nb - 1)
}