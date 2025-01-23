package piscine


func Join(strs []string, sep string) string {
	res := ""
	for i, str := range strs {
		if i!=0 {
			res += sep
		}
		res += str
	}
	return res
}