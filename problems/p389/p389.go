package p389

func findTheDifference(s string, t string) byte {
	sc := []byte(s)
	tc := []byte(t)

	buffer := make(map[byte]int32, len(sc))

	for _, c := range sc {
		buffer[c]++
	}

	for _, c := range tc {
		if buffer[c] == 0 {
			return c
		}
		buffer[c]--
	}
	return 0
}
