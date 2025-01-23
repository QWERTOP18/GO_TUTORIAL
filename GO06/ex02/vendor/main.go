package main

import (
	"os"
	"piscine"
)

// $ go mod init ex02
// $ go run .
// File name missing
// $ echo "Ultimate Question of Life, the Universe, and Everything" > 42.txt
// $ go run . 42.txt main.go
// Too many arguments
// $ go run . 42.txt
// Ultimate Question of Life, the Universe, and Everything



func main() {
	argc := ft.lenArray(os.Args)
	if argc == 1 {
		os.Stderr.Write([]byte("File name missing\n"))
		return
	}
	if argc > 2 {
		os.Stderr.Write([]byte("Too many arguments\n"))
		return
	}
	ft.printFile(os.Args[1])
}