package lc03

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	data := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
		"deefef":   2,
	}

	for s, expected := range data {
		got := lengthOfLongestSubstring(s)

		if got != expected {
			t.Fatalf("s: %v, got: %v, expected: %v", s, got, expected)
		}
	}
}
