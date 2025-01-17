package piscine

import "ft"

func PrintStr(s string) {
    for range s {
        ft.PrintRune(s)
    }
}