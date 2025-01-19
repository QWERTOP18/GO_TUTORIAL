package piscine

import "ft"

func IsPrime(nb int) int {
	if nb <= 1 {
		return 0
	}
	for i := 2; i < nb / i ; i++ {
		if nb % i == 0 {
			return 1
		} 
	}
	return 0
}