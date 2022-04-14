package p700

import "leet_code/models"

func searchBST(root *models.TreeNode, val int) *models.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if val > root.Val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}
