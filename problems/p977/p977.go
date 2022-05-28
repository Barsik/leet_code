package p977

func sortedSquares(nums []int) []int {
	l := len(nums)
	i, j, k := 0, l-1, l-1
	res := make([]int, l)
	for j >= i {
		if Abs(nums[j]) >= Abs(nums[i]) {
			res[k] = nums[j] * nums[j]
			j--
		} else {
			res[k] = nums[i] * nums[i]
			i++
		}
		k--
	}

	return res
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
