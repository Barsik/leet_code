package p1342

func numberOfSteps(num int) int {
	steps := 0
	if num == 0 {
		return 0
	}
	for num > 0 {
		if num&1 == 1 {
			steps += 2

		} else {
			steps++
		}
		num = num >> 1
	}
	return steps - 1
}
