package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
        return root
    } else if elem > root.Data {
        return BTreeSearchItem(root.Right, elem)
    } else  {
	    return BTreeSearchItem(root.Left, elem)
	}
	return nil
}
