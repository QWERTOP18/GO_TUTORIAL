package piscine


func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	for i:=1 ; i <= nb / i ; i++ {
		if i  == nb / i && nb % i == 0 {
			return i
		}
	}
	return 0
}