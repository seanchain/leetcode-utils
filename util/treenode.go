package util

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}

func createTreeNode(value int, left *TreeNode, right *TreeNode) *TreeNode {
	treeNode := new(TreeNode)
	treeNode.Val = value
	treeNode.left = left
	treeNode.right = right
	return treeNode
}

func (t *TreeNode) init(value int, left *TreeNode, right *TreeNode) {
	t.Val = value
	t.left = left
	t.right = right
}
