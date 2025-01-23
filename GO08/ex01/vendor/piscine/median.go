package piscine


func lenArr(s []int) (len int) {
	for range s {
		len++
	}
	return
}

func Median(s []int) int {
	n := lenArr(s)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	if n%2 == 0 {
		return (s[n/2-1] + s[n/2]) / 2
	}
	return s[n/2]
}