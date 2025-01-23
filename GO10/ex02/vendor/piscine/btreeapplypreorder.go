package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
    // Your implementation here
}
