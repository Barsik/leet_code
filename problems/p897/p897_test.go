package p897

import (
	"fmt"
	"leet_code/common"
	"leet_code/models"
	"testing"
)

var root = &models.TreeNode{Val: 5,
	Left:  &models.TreeNode{Val: 1},
	Right: &models.TreeNode{Val: 7}}

func TestIncreasingBST(t *testing.T) {
	result := increasingBST(root)
	arr := common.TreeNodeToArray(result)
	fmt.Printf("___ %v\n", arr)
	if !common.EqualArrays(arr, []int{1, 5, 7}) {
		t.Error("Expected [1,5,7]")
	}
}
