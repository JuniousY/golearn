package leetcode

func numSmallerByFrequency(queries []string, words []string) []int {
	counts := make([]int, 11)
	for _, word := range words {
		counts[countFor1170(word)]++
	}
	sum := 0
	for i, count := range counts {
		sum += count
		counts[i] = sum
	}
	results := make([]int, len(queries))
	for i, query := range queries {
		c := countFor1170(query)
		results[i] = counts[10] - counts[c]
	}
	return results
}

func countFor1170(word string) int {
	count := 0
	small := 'z'
	for _, c := range word {
		if c < small {
			small = c
			count = 1
		} else if c == small {
			count++
		}
	}
	return count
}
