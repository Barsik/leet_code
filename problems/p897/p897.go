package p897

import "leet_code/models"

var cur *models.TreeNode

func increasingBST(root *models.TreeNode) *models.TreeNode {
	ret := &models.TreeNode{}
	cur = ret
	increase(root)
	return ret.Right
}

func increase(node *models.TreeNode) {
	if node == nil {
		return
	}
	increase(node.Left)
	cur.Right = node
	node.Left = nil
	cur = node
	increase(node.Right)
}
