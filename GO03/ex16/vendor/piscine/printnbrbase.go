package piscine

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


func PrintNbrBase(nbr int, base string) {
	if base == "" {
		PrintNV();
		return
	}
	if !isBaseValid(base) {
		PrintNV();
		return
	}
	if nbr < 0 {
		fmt.Print("-")
		PrintNbrBase(-nbr, base)
		return
	}
	l := 0
	for range base {
		l++
	}


}
