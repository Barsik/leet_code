package p2011

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, o := range operations {

		x += int((rune(',') - rune(o[1])))

	}

	return x
}
