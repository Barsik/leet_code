package p3

func lengthOfLongestSubstring(s string) int {
	buffer := make(map[rune]int)
	arr := []rune(s)

	maxlen, curlen, i, slen := 0, 0, 0, len(arr)

	for i < slen {
		pos, ok := buffer[arr[i]]
		if !ok {
			curlen++
			buffer[arr[i]] = i
			i++
		} else {
			if maxlen < curlen {
				maxlen = curlen
			}
			buffer = make(map[rune]int)
			curlen = 0
			i = pos + 1
		}
	}

	if maxlen < curlen {
		maxlen = curlen
	}

	return maxlen
}
