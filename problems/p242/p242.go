package p242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := [26]int{}

	for i, _ := range s {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}

	for _, a := range arr {
		if a != 0 {
			return false
		}
	}

	return true
}
