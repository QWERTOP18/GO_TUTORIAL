package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func max(a,b int) int {
	if a < b {
		return b
    } 
	return a
}

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
        return 0
    }
    left := BTreeLevelCount(root.Left)
    right := BTreeLevelCount(root.Right)
    return max(left, right) + 1
}

