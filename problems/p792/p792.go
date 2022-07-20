package p792

func numMatchingSubseq(s string, words []string) (res int) {

	xwords := map[string]int{}
	for _, w := range words {
		xwords[w]++
	}

	for w, count := range xwords {

		i, j := 0, 0

		for i < len(s) && j < len(w) {
			if s[i] == w[j] {
				i++
				j++
			} else {
				i++
			}
		}

		if j == len(w) {
			res += count
		}

	}
	return
}
