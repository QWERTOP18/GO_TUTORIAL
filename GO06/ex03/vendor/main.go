
package main
import (
	"os"
	"piscine"
)

// $ echo "Born2Code" > quote.txt
// $ cat <<EOF> 42.txt
// "Ultimate Question of Life, the Universe, and Everything"
// EOF
// $ go mod init ex03
// $ go run . abc
// ERROR: open abc: no such file or directory
// exit status 1
// $ go run . quote.txt
// Born2Code
// $ go run . quote.txt abc
// Born2Code
// ERROR: open abc: No such file or directory
// $ cat quote.txt | ./cat
// Born2Code
// $ go run .
// Hello
// Hello
// ^C
// $ go run . quote.txt 42.txt
// Born2Code
// "Ultimate Question of Life, the Universe, and Everything"
// $

func main() {
	
	argc := piscine.LenArray(os.Args)
	if argc == 1 {
		piscine.PrintStdin()
	}
	for i := 1; i < argc; i++ {
		piscine.PrintFile(os.Args[i])
	}
}