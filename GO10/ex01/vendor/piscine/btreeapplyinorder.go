package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
    if root == nil {
		return 0, nil
	}
	left, _ := BTreeApplyInorder(root.Left, f)
	f(root.Data)
	right, _ = BTreeApplyInorder(root.Right, f)
	return left + right + 1, nil
}


