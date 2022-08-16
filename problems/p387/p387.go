package p387

func firstUniqChar(s string) int {
	buffer := [26]rune{}

	for _, c := range s {
		buffer[c-'a']++
	}

	for i, c := range s {
		if buffer[c-'a'] == 1 {
			return i
		}
	}

	return -1
}
