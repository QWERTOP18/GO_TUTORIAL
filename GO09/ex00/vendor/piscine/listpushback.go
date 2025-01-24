package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}


/**
*
* memoryのlifetimeがよくわからない
* Q 関数を抜けてもpopしないならmakeの必要性がない

* go build -gcflags="-m" 
*/
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
