package goevil

// levenshtein distance between two strings.
func distance(str1, str2 string) int {
	s1len := len(str1)
	s2len := len(str2)
	column := make([]int, len(str1)+1)

	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= s1len; y++ {
			oldkey := column[y]
			var incr int
			if str1[y-1] != str2[x-1] {
				incr = 1
			}

			column[y] = min(column[y]+1, min(column[y-1]+1, lastkey+incr))
			lastkey = oldkey
		}
	}
	return column[s1len]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
