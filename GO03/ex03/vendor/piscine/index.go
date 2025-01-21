package piscine


func Index(s string, toFind string) int {
	for i, v := range s {
		if string(v) == toFind {
			return i
		}
	}
	return -1
}
