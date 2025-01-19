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
    powpow := nb
    for i := 0; i < 32 && power > 0; i++ {
        if power & (1 << i) != 0 {
            res *= powpow
        }
        powpow *= powpow
    }
    return res
}
