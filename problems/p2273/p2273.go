package p2273

func removeAnagrams(words []string) []string {
	l := len(words)
	buffer := make(map[string][26]int, l)
	for i := l - 1; i >= 1; i-- {
		hash := getHash(words[i], buffer)
		phash := getHash(words[i-1], buffer)

		if isHashEquals(hash, phash) {
			words = append(words[:i], words[i+1:]...)
		}
	}

	return words
}

func isHashEquals(h1 [26]int, h2 [26]int) bool {
	for i := 0; i < 26; i++ {
		if h1[i] != h2[i] {
			return false
		}
	}
	return true
}

func getHash(str string, buffer map[string][26]int) [26]int {
	h, ok := buffer[str]

	if ok {
		return h
	}

	hash := [26]int{}
	for _, c := range str {
		hash[c-'a']++
	}

	buffer[str] = hash
	return hash
}
