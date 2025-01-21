package main
import (
	"os"
	"ft"
)

func PrintStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}
func main() {
	for _, s := range os.Args[1:] {
		PrintStr(s)
		ft.PrintRune('\n')
	}
}