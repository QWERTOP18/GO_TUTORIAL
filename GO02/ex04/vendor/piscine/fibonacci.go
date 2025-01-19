package piscine

import "ft"


func Fibonacci(index int) int {
	var dp[10000] int staticd;
	if index >= 10000 {
		return -1
	}
	if index < 2 {
		return 1
	}
	if dp[index] != 0 {
		return dp[index]
	}
	dp[index] = Fibonacci(index - 1) + Fibonacci(index - 2)
	return dp[index]
}