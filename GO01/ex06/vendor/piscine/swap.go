package piscine

import "ft"

func Swap(a *int, b *int) {
	a ^= b
	b ^= a
	a ^= b
}
