package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return 0, nil
	}
	f(root.Data)
	left, _ := BTreeApplyPreorder(root.Left, f)
	right, _ := BTreeApplyPreorder(root.Right, f)
	return left + right + 1, nil
}
