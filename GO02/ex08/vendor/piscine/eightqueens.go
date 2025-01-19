package piscine

import "ft"


func PrintLine(line [8]rune) {
	for _, r := range line {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func CheckDuplicate(depth int ,current [8]rune, next rune) bool{
	for i := 0; i < depth ; i++ {
		// column
		if current[i] == next {
			return true
		}
		// Diagonal
		if int(current[i])-int(next) == depth-i || int(current[i])-int(next) == i -depth{
            return true 
        }
	} 
	return false
}

func DepthFirstSearch(depth int, current [8]rune) {
	if depth == 8 {
		PrintLine(current)
		return 
	}
	for next:= '1'; next <= '8'; next++ {
		if !CheckDuplicate(depth, current, next) {
			current[depth] = next
			DepthFirstSearch(depth + 1, current)
		}
	}
}


func EightQueens() {
	// current := make([]rune, 8) 
	var current [8]rune
	DepthFirstSearch(0, current) 
}