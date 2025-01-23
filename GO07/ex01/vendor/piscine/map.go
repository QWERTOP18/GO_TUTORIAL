package piscine

func Map(f func(int) int, a []int) []bool {
	var res []bool
	for _, v := range a {
		res = append(res, f(v))
	}
	return res
}


func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n /i + 1; i++ {
		if n % i == 0 {
			return false
	}
	return true
}

