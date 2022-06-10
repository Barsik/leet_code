package p88

func merge(nums1 []int, m int, nums2 []int, n int) {
	l := len(nums1)
	buffer := make([]int, l)

	i, j, k := 0, 0, 0

	for k < l {
		if i < m && (j >= n || (nums1[i] <= nums2[j])) {
			buffer[k] = nums1[i]
			i++
		} else {
			buffer[k] = nums2[j]
			j++
		}
		k++
	}

	for w, b := range buffer {
		nums1[w] = b
	}
}
