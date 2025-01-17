package piscine

import "ft"

func UltimateDivMod(a *int, b *int) {
	if *b == 0 {
		//stderr
		fmt.Println("zero division")
		return
	}
	div := *a / *b
	mod := *a % *b
	*a = div
	*b = mod
}