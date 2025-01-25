package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
        return nil
    }
    for root.Right != nil {
        root = root.Right
    }
    return root
}
