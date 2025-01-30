package piscine

func Split(s, sep string) []string {
	var result []string
	for _, r := range


	
	return result
}


/* 
func SplitWhiteSpaces(str string) []string {
	var result []string
	var word string
	for _, s := range str {
		if s == ' ' || s == '\t' || s == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(s)
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
	 */