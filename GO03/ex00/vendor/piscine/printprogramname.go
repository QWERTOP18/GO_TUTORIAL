package piscine

import "ft"

func Printprogramname(nb int) int {
	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
	}
	return result
}
