package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node.Left == nil && node.Right == nil {
        return BTreeTransplant(root, node, nil)
    }
    if node.Left == nil {
        return BTreeTransplant(root, node, node.Right)
    }
    if node.Right == nil {
        return BTreeTransplant(root, node, node.Left)
    }
    min := BTreeMin(node.Right)
    node.Data = min.Data
    node.Right = BTreeDeleteNode(node.Right, min)
    return root
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


func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
        return nil
    }
    for root.Left != nil {
        root = root.Left
    }
    return root
}
