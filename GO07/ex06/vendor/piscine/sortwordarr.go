package piscine


func LenArr(a []string) (len int) {
	for range a {
		len++
	}
	return
}

func SortWordArr(a []string) {
	len := LenArr(a)
	for i := 0; i < len; i++ {
		for j := i; j < len; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}