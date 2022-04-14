package p700

import (
	"leet_code/common"
	"leet_code/models"
	"testing"
)

var root = &models.TreeNode{Val: 4,
	Left: &models.TreeNode{Val: 2,
		Left: &models.TreeNode{Val: 1,
			Left:  nil,
			Right: nil},
		Right: &models.TreeNode{Val: 3,
			Left:  nil,
			Right: nil}},
	Right: &models.TreeNode{Val: 7,
		Left:  nil,
		Right: nil}}

func TestSearchBST(t *testing.T) {
	result := searchBST(root, 2)
	arr := common.TreeNodeToArray(result)
	if !common.EqualArrays(arr, []int{2, 1, 3}) {
		t.Error("Expected [2,1,3]")
	}
}
