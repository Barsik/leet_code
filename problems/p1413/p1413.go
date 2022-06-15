package p1413

func minStartValue(nums []int) int {
	min := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if min > sum {
			min = sum
		}
	}

	if min >= 0 {
		return 1
	}

	return -1*min + 1
}
