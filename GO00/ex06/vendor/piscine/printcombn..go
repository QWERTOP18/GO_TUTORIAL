package piscine

import "ft"

func Rec(nbr [10]rune, depth int, maxDepth int) {
	if depth == maxDepth {
		for i := 1; i <= maxDepth; i++ {
			ft.PrintRune(nbr[i])
		}
		if nbr[1] != rune(10 - maxDepth) + '0'{
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
		return
	}

	for i := nbr[depth] + 1; i <= '9'; i++ {
		nbr[depth+1] = i
		Rec(nbr, depth+1, maxDepth)
	}
}

func PrintCombN(n int) {
	var nbr [10]rune

	for i := 0; i < n; i++ {
		nbr[i] = '0'
	}
	nbr[0] = '0' - 1
	Rec(nbr, 0, n)


	ft.PrintRune('\n')
}
