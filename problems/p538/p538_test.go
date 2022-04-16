package p538

import (
	"leet_code/common"
	"leet_code/models"
	"testing"
)

var root = &models.TreeNode{Val: 4,
	Left:  &models.TreeNode{Val: 1},
	Right: &models.TreeNode{Val: 6}}

func TestConvertBST(t *testing.T) {
	result := convertBST(root)
	arr := common.TreeNodeToArray(result)
	if !common.EqualArrays(arr, []int{10, 11, 6}) {
		t.Error("Expected [10,11,6]")
	}
}
