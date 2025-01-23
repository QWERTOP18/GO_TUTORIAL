package piscine

import "ft"

func lenArr(a []int) (len int) {
    for range a {
        len++
    }
    return 0
}

func Unmatch(a []int) int {
    len := lenArr(a)
    for i:= 0 ; i < len ; i++ {
        ok := false
        for j:= i + 1 ; j < len ; j++ {
            if a[i] == a[j] {
                ok = true
            }
            if !ok {
                return i
            }
        }
    }
    return -1
}
