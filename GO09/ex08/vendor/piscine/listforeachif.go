package piscine
type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func IsPositiveNode(node *NodeL) bool {
    switch node.Data.(type) {
    case int, float32, float64, byte:
        return node.Data.(int) > 0
    default:
        return false
    }
}

func IsAlNode(node *NodeL) bool {
    switch node.Data.(type) {
    case int, float32, float64, byte:
        return false
    default:
        return true
    }
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
    current := l.Head
    for current!= nil {
        if cond(current) {
            f(current)
        }
        current = current.Next
    }
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