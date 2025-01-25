package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
        return nil
    }
    for root.Left != nil {
        root = root.Left
    }
    return root
}
