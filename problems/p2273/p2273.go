package p2273

func removeAnagrams(words []string) []string {
	l := len(words)
	buffer := make(map[string]*[26]int, l)
	for i := l - 1; i >= 1; i-- {

		if isHashEquals(words[i], words[i-1], &buffer) {
			words = append(words[:i], words[i+1:]...)
		}
	}

	return words
}

func isHashEquals(str1 string, str2 string, buffer *map[string]*[26]int) bool {
	pH1 := getHash(str1, buffer)
	pH2 := getHash(str2, buffer)
	for i := 0; i < 26; i++ {
		if pH1[i] != pH2[i] {
			return false
		}
	}
	return true
}

func getHash(str string, buffer *map[string]*[26]int) *[26]int {
	ib := *buffer
	h, ok := ib[str]
	if ok {
		return h
	}
	hash := [26]int{}
	for _, c := range str {
		hash[c-'a']++
	}

	ib[str] = &hash
	return &hash
}
