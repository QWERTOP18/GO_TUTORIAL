package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListClear(l *List){
	for l.Head != nil {
        tmp := l.Head
        l.Head = l.Head.Next
        tmp.Next = nil
    }
    l.Tail = nil
}

