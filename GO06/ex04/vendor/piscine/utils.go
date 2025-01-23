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


func TailFile(fileName string, offset int, isHeader bool) {
	file, err := os.Open(fileName)
	if err != nil {
		os.Stderr.Write([]byte("tail: open " + fileName + ": no such file or directory\n"))
		return
	}
	defer file.Close()

	var buf [1024]byte
	var bufSize int
	var totalSize int
	for {
		n, err := file.Read(buf[:])
		if n > 0 {
			totalSize += n
			if totalSize > offset {
				if isHeader {
					os.Stdout.Write(buf[:n])
				} else {
					bufSize = n
				}
			}
		}
		if err != nil {
			break
		}
	}

	if !isHeader {
		os.Stdout.Write(buf[:bufSize])
	}
}


func stringToInt(s string) (int, bool) {
	var num int
	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i , _ := range s {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		if num > (int(^uint(0)>>1)-(int(s[i]-'0'))) / 10 {
			return 0, false
		}

		num = num*10 + int(s[i]-'0')
	}

	return num * sign, true
}