package piscine

import "ft"

func PrintComb() {
	for i := '0'; i <= '6' ; i++ {
		for j := i+1; j <= '8' ; j++ {
			for k := j+1; k <= '9' ; k++ {
				ft.PrintRune(i)
				ft.PrintRune(j)
				ft.PrintRune(k)
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
	ft.PrintRune('7')
	ft.PrintRune('8')
	ft.PrintRune('9')
	ft.PrintRune('\n')
}
