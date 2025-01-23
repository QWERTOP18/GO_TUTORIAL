import "ft"

const EvenMsg = "I have an even number of arguments\n"
const OddMsg = "I have an odd number of arguments\n"

type boolean int
const (
	no = iota
	yes
)

func printStr(s string) {
	for _, r := range s {
	ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}


func isEven(nbr int) boolean {
	if even(nbr) == 1 {
	return yes
	} else {
	return no
	}
}


func main() {
	if isEven(lengthOfArg) == 1 {
	printStr(EvenMsg)
	} else {
	printStr(OddMsg)
	}
}