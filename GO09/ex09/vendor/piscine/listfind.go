package piscine
type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func CompStr(a, b interface{}) bool {
    return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
    current := l.Head
    for current!= nil {
        if comp(current.Data, ref) {
            return &current.Data
        }
        current = current.Next
    }
    return nil
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