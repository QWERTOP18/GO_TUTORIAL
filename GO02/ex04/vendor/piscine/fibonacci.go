package piscine


// func Fibonacci(index int) int {
// 	var dp[10000] int static;
// 	if index >= 10000 {
// 		return -1
// 	}
// 	if index < 2 {
// 		return 1
// 	}
// 	if dp[index] != 0 {
// 		return dp[index]
// 	}
// 	dp[index] = Fibonacci(index - 1) + Fibonacci(index - 2)
// 	return dp[index]
// }

func Fibonacci(index int) int {
	if index < 3 {
		return 1
	}
	return Fibonacci(index - 1) + Fibonacci(index - 2)
}