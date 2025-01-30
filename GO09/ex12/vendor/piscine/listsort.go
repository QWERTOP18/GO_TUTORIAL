package piscine

type NodeI struct {
    Data int 
    Next *NodeI
}

func ListSort(l *NodeI) *NodeI{
	curr := l
	isSorted := true
	if l == nil {
		return nil
	}
	for curr.Next != nil {
		next := curr.Next
		if curr.Data > next.Data {
			isSorted = false
			curr.Data, next.Data = next.Data, curr.Data
		}
		curr = curr.Next
	}
	if isSorted {
		return l
	} else {
		return ListSort(l)
	}
}
