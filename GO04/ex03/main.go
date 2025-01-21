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


func StrCmp(a str, b str) int {
	
}

func Sort(argv []str, int len, _cmp func) {
	for i:=0; i<len -1 ; i++ {
		for j:=0; j<len ;j++ {
			if _cmp(argv[i], argv[j]) > 0 {
				argv[i], argv[j] = argv[j], argv[i]
			}
		}
	}
} 


func main() {
	len := CountLen(os.Args[1:])
	Sort(os.Argv[1:], len, StrCmp)
	for _, s := range os.Args[1:] {
		PrintStr(s)
		ft.PrintRune('\n')
	}
}
