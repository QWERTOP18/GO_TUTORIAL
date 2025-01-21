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
	lastSlash := 0
	for i, r := range os.Args[0] {
		if r == '/' {
			lastSlash = i
		}
	}
	PrintStr(os.Args[0][i:])
}
