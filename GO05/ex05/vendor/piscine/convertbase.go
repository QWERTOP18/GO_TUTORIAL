package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	deci := AtoiBase(nbr, baseFrom)
	return AtoiBase(deci, baseTo)
}


/*                    PREVIOUS EXERCISE                */ 


func isBaseValid(base string) bool {
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
	if !isBaseValid(base) {
		return 0
	}
	baseLen := 0
	for range base {
		baseLen++
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