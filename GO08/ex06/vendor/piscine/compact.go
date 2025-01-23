package piscine


func Compact(ptr *[]string) int {
	count := 0
	for _, str := range *ptr {
		if str != "" {
			count++
		}
	}
	res := make([]string, count)
	i := 0
	for _, str := range *ptr {
		if str != "" {
			res[i] = str
			i++
		}
	}
	*ptr = res
	return count
}
