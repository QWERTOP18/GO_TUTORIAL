package piscine

func isValidBase(base string) bool {
	l := 0
	for range base {
		l++
	}
	if l < 2 {
		return false
	}
	for i := 0; i < l-1; i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		for j := i + 1; j < l; j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}
	baseLen := 0
	for range base {
		baseLen++
	}
	if baseLen < 2 {
		return 0
	}
	res := 0
	for _, c := range s {
		index := 0
		for i, b := range base {
			if c == b {
				index = i
				break
			}
		}
		res = res*baseLen + index
	}
	return res
}