package p2325

import "strings"

func decodeMessage(key string, message string) string {
	var sb strings.Builder
	c := 'a'
	buffer := map[rune]rune{}

	for _, k := range key {
		if k != ' ' {
			_, ok := buffer[k]
			if !ok {
				buffer[k] = c
				c++
			}
		}
	}

	for _, m := range message {
		if m == ' ' {
			sb.WriteString(" ")
		} else {
			sb.WriteString(string(buffer[m]))
		}
	}

	return sb.String()

}
