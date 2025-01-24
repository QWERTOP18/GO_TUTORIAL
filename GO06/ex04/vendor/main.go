
package main
import (
	"os"
	"piscine"
)
func main() {
	
	argc := piscine.LenArray(os.Args)
	if argc == 1 {
		// piscine.PrintStdin()
		return 
	}
	if os.args[1] == "-c" {
		if argc == 2 {
			os.Stderr.Write([]byte("tail: option requires an argument -- c\n"))
			return 1
		}
		offset, ok := piscine.StringToInt(os.Args[2])
		if !ok {
			os.Stderr.Write([]byte("tail: illegal offset -- " + os.Args[2] + ":Invalid argument\n"))
			return 1
		}
		for i := 3; i < argc; i++ {
			piscine.TailFile(os.Args[i], offset, )
		}
	}
	for i := 1; i < argc; i++ {
		piscine.TailFile(os.Args[i])
	}
}