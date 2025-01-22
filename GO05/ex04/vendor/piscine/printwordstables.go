package piscine
import "ft"

func Println(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func PrintWordsTables(table []string) {
	for _, word := range table {
		Println(word)
	}
}