package piscine

import "ft"

const OPEN = true
const CLOSE = false

type Door struct {
	state bool
}
typedef ptrDoor *Door

func PrintStr(s string) {
for _, r := range s {
	ft.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	state = CLOSE
	return true
}

func IsDoorOpen(Door Door) {
	PrintStr("is the Door opened ?")
	return Door.state = OPEN
}
	
func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
}

func main() {
	door := &Door{}
	OpenDoor(door)
	if IsDoorClose(door) {
	OpenDoor(door)
	}
	if IsDoorOpen(door) {
	CloseDoor(door)
	}
	if door.state == OPEN {
	CloseDoor(door)
	}
}
