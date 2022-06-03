package p1480

func runningSum(nums []int) []int {
	sum := 0
	l := len(nums)

	for _, n := range nums {
		sum += n
	}

	for i := l - 1; i >= 0; i-- {
		temp := nums[i]
		nums[i] = sum
		sum -= temp
	}
	return nums
}
