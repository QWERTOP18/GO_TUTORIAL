package piscine

import "ft"

func StrLen(s string) int {
	count := 0
    for range s {
        count++
    }
    return count
}