package piscine

import "ft"

func Printf02(nb int)
{
	ft.PrintRune(nb / 10)
	ft.PrintRune(nb % 10)
}


func PrintComb2() {
	for i := 0; i < 100 ; i++ {
		for j := i+1; j < 100 ; j++ {
			prinff02(i)
			printf02(j)
		}
	}
	ft.PrintRune('\n')
}