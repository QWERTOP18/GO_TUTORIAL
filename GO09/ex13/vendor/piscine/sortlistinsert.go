package piscine


type NodeI struct {
    Data int 
    Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI{
    if l == nil {
        return &NodeI{Data: data_ref, Next: nil}
    }
    origin := l
    if l.Data > data_ref {
        return &NodeI{Data: data_ref, Next: l}
    }
    for l.Next != nil {
        next := l.Next
        if next.Data > data_ref {
            new := &NodeI{Data: data_ref, Next: next}
            l.Next = new
            return origin
        }
        l = next
    }
    l.Next = &NodeI{Data: data_ref, Next: nil}
    return origin
}