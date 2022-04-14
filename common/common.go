package common

import "leet_code/models"

func EqualArrays(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TreeNodeToArray(root *models.TreeNode) []int {
	var buffer []int
	if root != nil {
		buffer = append(buffer, root.Val)
	}
	treeNodeToArray(root, &buffer)
	return buffer
}

func treeNodeToArray(node *models.TreeNode, buffer *[]int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		*buffer = append(*buffer, node.Left.Val)
	}

	if node.Right != nil {
		*buffer = append(*buffer, node.Right.Val)
	}

	if node.Left != nil {
		treeNodeToArray(node.Left, buffer)
	}

	if node.Right != nil {
		treeNodeToArray(node.Right, buffer)
	}

	return
}
