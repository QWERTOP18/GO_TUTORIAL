package piscine

import (
	"fmt"
	"ft"
)

func DivMod(a int, b int, div *int, mod *int) {
	if b == 0 {
		//stderr
		fmt.Println("zero division")
		return
	}
	*div = a/b
	*mod = a%b
}