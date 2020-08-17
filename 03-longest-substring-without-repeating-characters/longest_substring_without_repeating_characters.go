package lc03

func lengthOfLongestSubstring(s string) int {
	result, l := 0, len(s)

	index := [128]int{}
	for i, j := 0, 0; j < l; j++ {
		if i < index[s[j]] {
			i = index[s[j]]
		}

		if result < j-i+1 {
			result = j - i + 1
		}
		index[s[j]] = j + 1
	}
	return result
}
