package p230

import "leet_code/models"

var kt int
var minVal int

func kthSmallest(root *models.TreeNode, k int) int {
	kt = k
	minVal = 0
	find(root)
	return minVal
}

func find(node *models.TreeNode) {
	if node == nil {
		return
	}

	find(node.Left)
	kt--
	if kt == 0 {
		minVal = node.Val
		return
	}

	find(node.Right)
}
