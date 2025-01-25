package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}

// if BTreeIsBinary(root.Left) == false {
// 	return false
// }
// if BTreeIsBinary(root.Right) == false {
// 	return false
// }

