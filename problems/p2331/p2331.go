package p2331

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if root == nil {
		return false
	} else if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	} else if root.Val == 3 {
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	} else {
		return root.Val == 1
	}
}
