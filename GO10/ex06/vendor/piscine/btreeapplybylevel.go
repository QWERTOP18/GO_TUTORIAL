package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
    // Your implementation here
}

