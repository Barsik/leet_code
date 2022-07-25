package p34

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	n := len(nums)
	if n == 0 {
		return ret
	}
	i, j := 0, n-1

	for i < j {
		mid := (i + j) / 2

		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid
		}
	}

	if nums[i] != target {
		return ret
	}

	ret[0] = i

	j = n - 1

	for i < j {
		mid := (i + j + 1) / 2

		if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid
		}
	}

	ret[1] = j

	return ret
}
