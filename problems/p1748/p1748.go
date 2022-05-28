package p1748

func sumOfUnique(nums []int) int {
	buffer := [101]int{}

	for _, n := range nums {
		buffer[n]++
	}
	sum := 0
	for i, b := range buffer {
		if b == 1 {
			sum += i
		}
	}

	return sum
}
