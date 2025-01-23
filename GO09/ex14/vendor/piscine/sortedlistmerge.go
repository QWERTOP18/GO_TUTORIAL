package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func SortedListMerge(n1 *NodeL, n2 *NodeL) *NodeL {
    // Your implementation here
}

