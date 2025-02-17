package piscine
type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func ListMerge(l1 *List, l2 *List) {
    if l1.Head == nil {
        l1.Head = l2.Head
        l1.Tail = l2.Tail
        return
    }
    l1.Tail.Next = l2.Head
    l1.Tail = l2.Tail
}

/******************* PREVIOUS EXERCISE *******************/

func ListPushBack(l *List, data interface{}) {
	node := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
	
}