package piscine


func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i <= nb / i ; i++ {
		if nb % i == 0 {
			return false
		} 
	}
	return true
}

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	}
	return FindNextPrime(nb + 1)
}