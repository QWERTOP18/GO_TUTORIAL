package piscine

import "ft"

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

func Putnbr(n int) {
	if n >= 10 || n <= -10 {
		Putnbr(n / 10)
	}
	ft.PrintRune(rune(n % 10 + '0'))
}

func PrintNbr(n int) {
	if n == 0 {
		ft.PrintRune('0')
		ft.PrintRune('\n')
		return 
	} 
	if n < 0 {
		ft.PrintRune('-') 
	}
	ft.Putnbr(n)
	ft.PrintRune('\n')
}