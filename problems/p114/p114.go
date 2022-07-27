package p114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	cur := root

	for cur != nil {

		if cur.Left != nil {
			temp := cur.Left

			for temp.Right != nil {
				temp = temp.Right
			}

			temp.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}

		cur = cur.Right
	}
}
