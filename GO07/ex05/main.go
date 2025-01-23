package main

import (
	"fmt"
	"os"
)

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

func add(a, b int) {
	if (b > 0 && a > (int(^uint(0)>>1)-b)) || (b < 0 && a < (-int(^uint(0)>>1)-1-b)) {
		return
	}
	fmt.Println(a + b)
}

func sub(a, b int) {
	if (b < 0 && a > (int(^uint(0)>>1)+b)) || (b > 0 && a < (-int(^uint(0)>>1)-1+b)) {
		return
	}
	fmt.Println(a - b)
}

func mul(a, b int) {
	if b != 0 && (a > (int(^uint(0)>>1)/b) || a < (-int(^uint(0)>>1)-1)/b) {
		return
	}
	fmt.Println(a * b)
}

func div(a, b int) {
	if b == 0 {
		fmt.Println("No division by 0")
		return
	}
	fmt.Println(a / b)
}

func mod(a, b int) {
	if b == 0 {
		fmt.Println("No modulo by 0")
		return
	}
	fmt.Println(a % b)
}

func LenArr(a []string) (len int) {
	for range a {
		len++
	}
	return
}

func main() {
	if LenArr(os.Args) != 4 {
		return
	}

	num1, ok1 := stringToInt(os.Args[1])
	num2, ok2 := stringToInt(os.Args[3])
	op := os.Args[2]


	switch op {
	case "+":
		add(num1, num2)
	case "-":
		sub(num1, num2)
	case "*":
		mul(num1, num2)
	case "/":
		div(num1, num2)
	case "%":
		mod(num1, num2)
	default:
		// fmt.Println("Unsupported operator. Use +, -, *, /, or %.")
	}
}
