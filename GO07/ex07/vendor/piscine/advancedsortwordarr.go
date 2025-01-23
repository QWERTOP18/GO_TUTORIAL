package piscine

func LenArr(a []string) (len int) {
	for range a {
		len++
	}
	return
}

func Compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func AdvancedSortWordArr(a []string, f func(a,b string) int) {
	len := LenArr(a)
	for i := 0; i < len; i++ {
		for j := i; j < len; j++ {
			if f(a[i], a[j]) > 0 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}