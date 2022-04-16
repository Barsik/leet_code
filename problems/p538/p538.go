package p538

import "leet_code/models"

var sum int

func convertBST(root *models.TreeNode) *models.TreeNode {
	sum = 0
	convert(root)
	return root
}

func convert(node *models.TreeNode) {
	if node == nil {
		return
	}
	if node.Right != nil {
		convert(node.Right)
	}
	sum += node.Val
	node.Val = sum

	if node.Left != nil {
		convert(node.Left)
	}
}
