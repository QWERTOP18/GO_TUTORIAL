package piscine
import "ft"

func Println(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func PrintWordsTables(table []string) {
	for _, word := range table {
		Println(word)
	}
}

/*                    PREVIOUS EXERCISE                */ 

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
	