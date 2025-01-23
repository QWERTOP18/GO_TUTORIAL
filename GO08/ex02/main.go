package main

import (
	"os"
)

const NGWORDS = ["42", "piscine", "piscine 42"]


func main() {
	for _, arg := range os.Args[1:] {
		if arg in NGWORDS {
			os.Stderr.WriteString("Alert!!!\n")
		}
	}
}