package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
    if root == nil {
		return
	}
	var queue [100]*TreeNode
	front := 0
	rear := 0

	queue[rear] = root
	rear++

	for front < rear {
		node := queue[front]
		front++

		_, err := f(node.Data)
		if err != nil {
			return
		}

		if node.Left != nil {
			queue[rear] = node.Left
			rear++
		}

		if node.Right != nil {
			queue[rear] = node.Right
			rear++
		}
}

