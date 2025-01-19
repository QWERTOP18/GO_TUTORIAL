package piscine

// func IterativePower(nb int, power int) int {
// 	res := 1
// 	for i := 1; i <= power; i++ {
// 		res *= i
// 	} 
// 	return res
// }
func IterativePower(nb int, power int) int {
	res := 1
	powpow = power
	for i := 0; i <= (power >> i); i++ {
		if (power & 1 << i) {
			res *= i
		}
	} 
	return res
}