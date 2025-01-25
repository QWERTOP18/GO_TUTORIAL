package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
        return rplc
    }
    if root.Left == node {
        root.Left = rplc
    } else if root.Right == node {
        root.Right = rplc
    }
    if rplc != nil {
        rplc.Parent = root.Parent
    }
    if root.Parent == nil {
        return rplc
    }
    if root.Parent.Left == root {
        root.Parent.Left = rplc
    } else {
        root.Parent.Right = rplc
    }
    return root
}
