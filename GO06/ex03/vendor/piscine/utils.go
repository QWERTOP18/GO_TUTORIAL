package piscine
import (
	"os"
)


func lenArray(args []string) (len int) {
	for range args {
		len++
	}
	return
}

func printBuffer(buf []byte) {
	os.Stdout.Write(buf)
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		os.Stderr.Write([]byte("open " + fileName + ": no such file or directory\n"))
		return
	}
	defer file.Close()

	var buf [1024]byte
	for {
		n, err := file.Read(buf[:])
		if n > 0 {
			os.Stdout.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}
}

func PrintStdin() {
	var buf [1024]byte
	for {
		n, err := os.Stdin.Read(buf[:])
		if n > 0 {
			os.Stdout.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}
}
