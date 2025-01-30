package piscine

type NodeI struct {
    Data int 
    Next *NodeI
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	n1last := n1
	for n1last.Next != nil {
		n1last = n1last.Next
	}
	n1last.Next = n2
	return ListSort(n1)
}

/******************* PREVIOUS EXERCISE *******************/
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
