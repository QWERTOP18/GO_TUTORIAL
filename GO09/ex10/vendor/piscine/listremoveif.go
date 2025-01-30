package piscine
type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
    current := l.Head
    var prev *NodeL
    
    for current != nil {
        if current.Data == data_ref {
            if current == l.Head {
                l.Head = current.Next
                if l.Head == nil {
                    l.Tail = nil
                }
            } else {
                prev.Next = current.Next
                if current.Next == nil {
                    l.Tail = prev
                }
            }
            current = current.Next
            continue
        }
        
        prev = current
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