package p2248

func intersection(nums [][]int) []int {
	counter := [1001]int{}
	rows := len(nums)
	for _, arr := range nums {

		for _, a := range arr {
			counter[a]++
		}
	}

	res := make([]int, 0, 1001)

	for i, c := range counter {
		if c == rows {
			res = append(res, i)
		}
	}
	return res
}
