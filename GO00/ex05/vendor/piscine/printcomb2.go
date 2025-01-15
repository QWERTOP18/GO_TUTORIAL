package piscine

import "ft"

func Printf02(nb int) {
	ft.PrintRune(rune(nb / 10) + '0')
	ft.PrintRune(rune(nb % 10) + '0')
}


func PrintComb2() {
	for i := 0; i < 100 ; i++ {
		for j := i+1; j < 100 ; j++ {
			Printf02(i)
			ft.PrintRune(' ')
			Printf02(j)
			if !(i == 98 && j == 99) {
                ft.PrintRune(',')
                ft.PrintRune(' ')
            }
		}
	}
	ft.PrintRune('\n')
}