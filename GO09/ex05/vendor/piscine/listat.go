type NodeL struct {
    Data interface{}
    Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
    if pos < 0 || l == nil {
        return nil
    }

    current := l.Head
    for i := 0; i < pos && current!= nil; i++ {
        current = current.Next
    }

    return current
}
