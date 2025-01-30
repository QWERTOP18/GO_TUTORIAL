package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	ct := 0
	it := l.Head
	for it != nil {
        ct++
        it = it.Next
    }
	return ct
}


/******************* PREVIOUS EXERCISE *******************/

func ListPushFront(l *List, data interface{}) {
	node := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
	
}