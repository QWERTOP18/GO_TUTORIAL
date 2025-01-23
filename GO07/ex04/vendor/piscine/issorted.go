package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	length := 0
	isAsc := true
	isDesc := true
	for range tab {
		length++
	}
	if length < 2 {
		return true
	}
	for i := 1; i < length && (isAsc || isDesc); i++ {
		if f(tab[i-1], tab[i]) > 0 {
			isAsc = false
		}
		if f(tab[i-1], tab[i]) < 0 {
			isDesc = false
		}
	}
	return isAsc || isDesc
}