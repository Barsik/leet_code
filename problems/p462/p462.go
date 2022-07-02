package p462

import "sort"

func minMoves2(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}
	sort.Ints(nums)
	mi := l / 2
	res := 0
	for i, n := range nums {
		if i < mi {
			res += (nums[mi] - n)
		} else if mi < i {
			res += (n - nums[mi])
		}
	}
	return res
}
