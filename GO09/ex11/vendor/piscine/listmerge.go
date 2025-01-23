type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func ListMerge(l1 *List, l2 *List) {
    // Your implementation here
}

