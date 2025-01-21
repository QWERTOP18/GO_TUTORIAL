package main
import (
	"os"
	"ft"
)

func CountLen(argv []string) (len int) {
	for range argv {
		len ++
	}
	return 
}


func PrintStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func main() {
	len := CountLen(os.Args[1:])
	for i:=len ; i >= 1 ; i--  {
		PrintStr(os.Args[i])
		ft.PrintRune('\n')
	}
}