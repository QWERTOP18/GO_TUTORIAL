package main
import (
	"fmt"
	"piscine"
)
func main() {
	fmt.Println(piscine.IsPrime(5))
	fmt.Println(piscine.IsPrime(4))
	fmt.Println(piscine.IsPrime(2147483647))
	fmt.Println(piscine.IsPrime(9223372036854775807))

}