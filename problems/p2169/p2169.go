package p2169

func countOperations(num1 int, num2 int) int {
	count := 0
	for num1 > 0 && num2 > 0 {
		if num2 > num1 {
			num1, num2 = num2, num1
		}
		count += num1 / num2

		num1 %= num2
	}
	return count
}
