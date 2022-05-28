package p268

func missingNumber(nums []int) int {
	sum := 0
	i := 0

	for ; i < len(nums); i++ {
		sum = sum + i - nums[i]
	}

	return sum + i
}
