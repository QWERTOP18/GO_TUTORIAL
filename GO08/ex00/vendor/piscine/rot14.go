package piscine

func Rot14(s string) string {
	var res string
	for _, r := range s {
		if 'A' <= r && r <= 'Z' {
			res += string('A' + (r-'A'+14)%26)
		} else if 'a' <= r && r <= 'z' {
			res += string('a' + (r-'a'+14)%26) 
		} else {
			res += string(r)
		}
	}
	return res
}

